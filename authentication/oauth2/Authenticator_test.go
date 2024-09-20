package oauth2

import (
	"net/url"
	"strings"
	"testing"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/authentication/oauth2/errors"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/configuration"
)

var accessTokenResponseSuccess = `{
  "access_token": "accessToken",
  "expires_in": 300
}`

var accessTokenResponseInvalidClient = `{
  "error": "unauthorized_client",
  "error_description": "INVALID_CREDENTIALS: Invalid client credentials"
}`

var accessTokenResponseExpired = `{
  "access_token": "expiredAccessToken",
  "expires_in": -1
}`

type mockConnection struct {
	statusCode      int
	responseHeaders []communication.Header
	responseBody    string
	callCount       int
}

func (c *mockConnection) Post(resourceURI url.URL, requestHeaders []communication.Header, body string, handler communication.ResponseHandler) (interface{}, error) {
	c.callCount++
	return handler(c.statusCode, c.responseHeaders, strings.NewReader(c.responseBody))
}
func (c *mockConnection) Close() error {
	return nil
}

type oauth2Result struct {
	authorization string
	err           error
}

func TestSuccessfulOAuth2Authentication(t *testing.T) {
	contentType, _ := communication.NewHeader("Content-Type", "application/json")
	c := &mockConnection{
		statusCode:      200,
		responseHeaders: []communication.Header{*contentType},
		responseBody:    accessTokenResponseSuccess,
	}

	config, err := configuration.DefaultOAuth2Configuration("clientId", "clientSecret", "http://localhost/auth/realms/api/protocol/openid-connect/token", "Test")
	if err != nil {
		t.Fatalf("TestSuccessfulOAuth2Authentication : %v", err)
	}

	authenticator, err := newAuthenticator(config, func() (oauth2Connection, error) {
		return c, nil
	})
	if err != nil {
		t.Fatalf("TestSuccessfulOAuth2Authentication : %v", err)
	}

	results := make(chan oauth2Result, 10)

	for i := 0; i < 10; i++ {
		go func() {
			authorization, err := authenticator.GetAuthorization("", url.URL{Path: "/api/v1/payments"}, nil)
			results <- oauth2Result{authorization, err}
		}()
	}

	for i := 0; i < 10; i++ {
		result := <-results
		expectedAuthorization := "Bearer accessToken"
		if result.authorization != expectedAuthorization {
			t.Fatalf("TestSuccessfulOAuth2Authentication : expectedAuthorization '%s' != '%s'", expectedAuthorization, result.authorization)
		}
		if result.err != nil {
			t.Fatalf("TestSuccessfulOAuth2Authentication : %v", err)
		}
	}

	if c.callCount != 1 {
		t.Fatalf("TestSuccessfulOAuth2Authentication : expectedCallCount '1' != '%d'", c.callCount)
	}
}

func TestFailedOAuth2Authentication(t *testing.T) {
	contentType, _ := communication.NewHeader("Content-Type", "application/json")
	c := &mockConnection{
		statusCode:      401,
		responseHeaders: []communication.Header{*contentType},
		responseBody:    accessTokenResponseInvalidClient,
	}

	config, err := configuration.DefaultOAuth2Configuration("clientId", "clientSecret", "http://localhost/auth/realms/api/protocol/openid-connect/token", "Test")
	if err != nil {
		t.Fatalf("TestFailedOAuth2Authentication : %v", err)
	}

	authenticator, err := newAuthenticator(config, func() (oauth2Connection, error) {
		return c, nil
	})
	if err != nil {
		t.Fatalf("TestFailedOAuth2Authentication : %v", err)
	}

	results := make(chan oauth2Result, 10)

	for i := 0; i < 10; i++ {
		go func() {
			authorization, err := authenticator.GetAuthorization("", url.URL{Path: "/api/v1/payments"}, nil)
			results <- oauth2Result{authorization, err}
		}()
	}

	for i := 0; i < 10; i++ {
		result := <-results
		switch oe := result.err.(type) {
		case *errors.OAuth2Error:
			{
				expectedError := "There was an error while retrieving the OAuth2 access token: unauthorized_client - INVALID_CREDENTIALS: Invalid client credentials"
				if oe.Error() != expectedError {
					t.Fatalf("TestFailedOAuth2Authentication : expectedError '%s' != '%s'", expectedError, oe.Error())
				}

				break
			}
		default:
			{
				t.Fatalf("TestFailedOAuth2Authentication : %v", result.err)
			}
		}
		if result.authorization != "" {
			t.Fatalf("TestFailedOAuth2Authentication : expectedAuthorization '%s' != '%s'", "", result.authorization)
		}
	}

	if c.callCount != 10 {
		t.Fatalf("TestFailedOAuth2Authentication : expectedCallCount '10' != '%d'", c.callCount)
	}
}

func TestExpiredOAuth2Authentication(t *testing.T) {
	contentType, _ := communication.NewHeader("Content-Type", "application/json")
	c := &mockConnection{
		statusCode:      200,
		responseHeaders: []communication.Header{*contentType},
		responseBody:    accessTokenResponseExpired,
	}

	config, err := configuration.DefaultOAuth2Configuration("clientId", "clientSecret", "http://localhost/auth/realms/api/protocol/openid-connect/token", "Test")
	if err != nil {
		t.Fatalf("TestExpiredOAuth2Authentication : %v", err)
	}

	authenticator, err := newAuthenticator(config, func() (oauth2Connection, error) {
		return c, nil
	})
	if err != nil {
		t.Fatalf("TestExpiredOAuth2Authentication : %v", err)
	}

	results := make(chan oauth2Result, 10)

	for i := 0; i < 10; i++ {
		go func() {
			authorization, err := authenticator.GetAuthorization("", url.URL{Path: "/api/v1/payments"}, nil)
			results <- oauth2Result{authorization, err}
		}()
	}

	for i := 0; i < 10; i++ {
		result := <-results
		expectedAuthorization := "Bearer expiredAccessToken"
		if result.authorization != expectedAuthorization {
			t.Fatalf("TestExpiredOAuth2Authentication : expectedAuthorization '%s' != '%s'", expectedAuthorization, result.authorization)
		}
		if result.err != nil {
			t.Fatalf("TestExpiredOAuth2Authentication : %v", err)
		}
	}

	if c.callCount != 10 {
		t.Fatalf("TestExpiredOAuth2Authentication : expectedCallCount '10' != '%d'", c.callCount)
	}
}
