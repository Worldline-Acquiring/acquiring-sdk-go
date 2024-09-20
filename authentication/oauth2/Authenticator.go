package oauth2

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"strings"
	"sync"
	"time"

	oauth2Errors "github.com/Worldline-Acquiring/acquiring-sdk-go/authentication/oauth2/errors"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/configuration"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/json"
)

type oauth2Connection interface {
	Post(resourceURI url.URL, requestHeaders []communication.Header, body string, handler communication.ResponseHandler) (interface{}, error)
	Close() error
}

type accessToken struct {
	accessToken    string
	expirationTime time.Time
}

type accessTokenResponse struct {
	AccessToken      *string        `json:"access_token"`
	ExpiresIn        *time.Duration `json:"expires_in"`
	Error            *string        `json:"error"`
	ErrorDescription *string        `json:"error_description"`
}

type tokenType struct {
	path        string
	scopes      string
	accessToken *accessToken
	lock        sync.RWMutex
}

func newTokenType(path, tokens string) *tokenType {
	return &tokenType{path, tokens, nil, sync.RWMutex{}}
}

func (t *tokenType) isAccessTokenNullOrExpired() bool {
	return t.accessToken == nil || t.accessToken.expirationTime.Before(time.Now())
}

// Authenticator represents an authentication.Authenticator implementation using OAuth2
type Authenticator struct {
	clientID          string
	clientSecret      string
	tokenURI          url.URL
	connectionFactory func() (oauth2Connection, error)
	marshaller        json.Marshaller
	accessTokens      []*tokenType
}

// NewAuthenticator creates an OAuth2Authenticator using values from the given CommunicatorConfiguration.
// Its OAuth2 client id, client secret and token URI must not be empty, and its connect timeout and socket timeout must be positive.
func NewAuthenticator(conf *configuration.CommunicatorConfiguration) (*Authenticator, error) {
	// Capture current values
	connectTimeout := conf.ConnectTimeout
	socketTimeout := conf.SocketTimeout
	keepAliveTimeout := time.Duration(0)
	idleTimeout := time.Duration(0)
	proxy := conf.Proxy

	connectionFactory := func() (oauth2Connection, error) {
		return communicator.NewDefaultConnection(socketTimeout, connectTimeout, keepAliveTimeout, idleTimeout, 1, proxy)
	}

	return newAuthenticator(conf, connectionFactory)
}

func newAuthenticator(conf *configuration.CommunicatorConfiguration, connectionFactory func() (oauth2Connection, error)) (*Authenticator, error) {
	if strings.TrimSpace(conf.GetOAuth2ClientID()) == "" {
		return nil, errors.New("OAuth2ClientID is required")
	}
	if strings.TrimSpace(conf.GetOAuth2ClientSecret()) == "" {
		return nil, errors.New("OAuth2ClientSecret is required")
	}
	if strings.TrimSpace(conf.OAuth2TokenURI) == "" {
		return nil, errors.New("OAuth2TokenURI is required")
	}
	if conf.ConnectTimeout.Nanoseconds() <= 0 {
		return nil, errors.New("ConnectTimeout must be positive")
	}
	if conf.SocketTimeout.Nanoseconds() <= 0 {
		return nil, errors.New("SocketTimeout must be positive")
	}

	tokenURI, err := url.Parse(conf.OAuth2TokenURI)
	if err != nil {
		return nil, err
	}

	marshaller := json.DefaultMarshaller()

	// Only a limited amount of scopes may be sent in one request.
	// While at the moment all scopes fit in one request, keep this code so we can easily add more token types if necessary.
	// The empty path will ensure that all paths will match, as each full path ends with an empty string.
	accessTokens := []*tokenType{
		newTokenType("", "processing_payment processing_refund processing_credittransfer "+
			"processing_accountverification processing_operation_reverse processing_dcc_rate services_ping"),
	}

	authenticator := Authenticator{
		clientID:          conf.GetOAuth2ClientID(),
		clientSecret:      conf.GetOAuth2ClientSecret(),
		tokenURI:          *tokenURI,
		connectionFactory: connectionFactory,
		marshaller:        marshaller,
		accessTokens:      accessTokens,
	}
	return &authenticator, nil
}

// GetAuthorization returns an OAuth2 bearer token including the Bearer prefix
func (a *Authenticator) GetAuthorization(httpMethod string, resourceURI url.URL, requestHeaders []communication.Header) (string, error) {
	t, err := a.getTokenType(resourceURI.Path)
	if err != nil {
		return "", err
	}
	t.lock.RLock()
	if t.isAccessTokenNullOrExpired() {
		// Replace the read lock with a write lock
		t.lock.RUnlock()
		t.lock.Lock()
		// Check again, it may have become false because of another access token call
		if t.isAccessTokenNullOrExpired() {
			token, err := a.getAccessToken(t.scopes)
			if err != nil {
				t.lock.Unlock()
				return "", err
			}
			t.accessToken = token
		}
		// Replace the write lock with a read lock again
		t.lock.Unlock()
		t.lock.RLock()
	}

	authorization := "Bearer " + t.accessToken.accessToken

	t.lock.RUnlock()

	return authorization, nil
}

func (a *Authenticator) getTokenType(fullPath string) (*tokenType, error) {
	for _, t := range a.accessTokens {
		if strings.HasSuffix(fullPath, t.path) || strings.Contains(fullPath, t.path+"/") {
			return t, nil
		}
	}
	return nil, oauth2Errors.NewOAuth2Error("Scope could not be found for path " + fullPath)
}

func (a *Authenticator) getAccessToken(scopes string) (*accessToken, error) {
	header, err := communication.NewHeader("Content-Type", "application/x-www-form-urlencoded")
	if err != nil {
		return nil, err
	}
	requestHeaders := []communication.Header{*header}
	requestBody := fmt.Sprintf("grant_type=client_credentials&client_id=%s&client_secret=%s&scope=%s", a.clientID, a.clientSecret, scopes)

	connection, err := a.connectionFactory()
	if err != nil {
		return nil, err
	}
	defer func(connection oauth2Connection) {
		_ = connection.Close()
	}(connection)

	startTime := time.Now()

	var accessToken accessToken

	_, err = connection.Post(a.tokenURI, requestHeaders, requestBody, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
		var accessTokenResponse accessTokenResponse

		err := a.marshaller.UnmarshalFromReader(reader, &accessTokenResponse)
		if err != nil {
			return nil, err
		}

		if statusCode != 200 {
			return nil, oauth2Errors.NewOAuth2Error(fmt.Sprintf("There was an error while retrieving the OAuth2 access token: %s - %s",
				*accessTokenResponse.Error, *accessTokenResponse.ErrorDescription,
			))
		}

		expiresIn := *accessTokenResponse.ExpiresIn * time.Second
		accessToken.expirationTime = startTime.Add(expiresIn)
		accessToken.accessToken = *accessTokenResponse.AccessToken

		return nil, nil
	})
	if err != nil {
		return nil, err
	}

	return &accessToken, nil
}
