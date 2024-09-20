package configuration

import (
	"net/url"
	"time"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/domain"
)

// CommunicatorConfiguration represents the configuration to be used by a Communicator
type CommunicatorConfiguration struct {
	// APIEndpoint represents the API endpoint of for the communicator
	APIEndpoint url.URL
	// ConnectTimeout represents the connect timeout
	ConnectTimeout time.Duration
	// SocketTimeout represents the request timeout
	SocketTimeout time.Duration
	// IdleTimeout represents the idle connection timeout
	IdleTimeout time.Duration
	// KeepAliveTimeout represents the HTTP KeepAlive interval
	KeepAliveTimeout time.Duration
	// MaxConnections represents the maximum amount of concurrent pooled connections
	MaxConnections int
	// AuthorizationType represents authorizationType used to sign the requests
	AuthorizationType AuthorizationType
	// AuthorizationID represents an id used for authorization. The meaning of this id is different for each authorization type.
	// For instance, for OAuth2 this is the client ID.
	AuthorizationID string
	// AuthorizationSecret represents a secret used for authorization. The meaning of this secret is different for each authorization type.
	// For instance, for OAuth2 this is the client secret.
	AuthorizationSecret string
	// OAuth2TokenURI represents the OAuth2 token URI.
	OAuth2TokenURI string
	// Proxy represents the URL for the connection proxy
	Proxy *url.URL
	// Integrator represents the integrator name
	Integrator string
	// ShoppingCartExtension represents the shopping cart extension used in the metadata headers
	ShoppingCartExtension *domain.ShoppingCartExtension
}

// GetOAuth2ClientID returns the OAuth2 client ID.
//
// This function is an alias for getting c.AuthorizationID
func (c *CommunicatorConfiguration) GetOAuth2ClientID() string {
	return c.AuthorizationID
}

// SetOAuth2ClientID sets the OAuth2 client ID.
//
// This function is an alias for setting c.AuthorizationID
func (c *CommunicatorConfiguration) SetOAuth2ClientID(oauth2ClientID string) {
	c.AuthorizationID = oauth2ClientID
}

// GetOAuth2ClientSecret returns the OAuth2 client secret.
//
// This function is an alias for getting c.AuthorizationSecret
func (c *CommunicatorConfiguration) GetOAuth2ClientSecret() string {
	return c.AuthorizationSecret
}

// SetOAuth2ClientSecret sets the OAuth2 client secret.
//
// This function is an alias for setting c.AuthorizationSecret
func (c *CommunicatorConfiguration) SetOAuth2ClientSecret(oauth2ClientSecret string) {
	c.AuthorizationSecret = oauth2ClientSecret
}

// The default configuration used by the factory is the following:
// APIEndpoint: api.preprod.acquiring.worldline-solutions.com
// ConnectTimeout: 5 seconds
// SocketTimeout: 30 seconds
// IdleTimeout: 5 seconds
// KeepAliveTimeout: 30 seconds
// MaxConnections: 10
// Proxy: none
var defaultConfiguration = CommunicatorConfiguration{
	APIEndpoint: url.URL{
		Scheme: "https",
		Host:   "api.preprod.acquiring.worldline-solutions.com",
	},
	ConnectTimeout:   5 * time.Second,
	SocketTimeout:    30 * time.Second,
	IdleTimeout:      5 * time.Second,
	KeepAliveTimeout: 30 * time.Second,
	MaxConnections:   10,
}

// DefaultOAuth2Configuration returns the default communicator configuration for authenticating using OAuth2
func DefaultOAuth2Configuration(clientID, clientSecret, tokenURI, integrator string) (*CommunicatorConfiguration, error) {
	customizedConfiguration := defaultConfiguration

	customizedConfiguration.Integrator = integrator
	customizedConfiguration.AuthorizationType = OAUTH2
	customizedConfiguration.AuthorizationID = clientID
	customizedConfiguration.AuthorizationSecret = clientSecret
	customizedConfiguration.OAuth2TokenURI = tokenURI

	return &customizedConfiguration, nil
}
