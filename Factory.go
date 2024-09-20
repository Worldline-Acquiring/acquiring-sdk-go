package acquiringsdk

import (
	"errors"
	"net/url"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/authentication"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/authentication/oauth2"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/configuration"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/json"
)

// CreateOAuth2Configuration creates a CommunicatorConfiguration with default OAuth2 settings and the given clientID, clientSecret and tokenURI
func CreateOAuth2Configuration(clientID, clientSecret, tokenURI, integrator string) (*configuration.CommunicatorConfiguration, error) {
	return configuration.DefaultOAuth2Configuration(clientID, clientSecret, tokenURI, integrator)
}

// CreateCommunicatorBuilderFromConfiguration creates a CommunicatorBuilder with the given CommunicatorConfiguration
func CreateCommunicatorBuilderFromConfiguration(config *configuration.CommunicatorConfiguration) (*CommunicatorBuilder, error) {
	builder := NewCommunicatorBuilder()

	connection, err := communicator.NewDefaultConnection(config.SocketTimeout,
		config.ConnectTimeout,
		config.KeepAliveTimeout,
		config.IdleTimeout,
		config.MaxConnections,
		config.Proxy)
	if err != nil {
		return nil, err
	}

	metadataProviderBuilder, err := communicator.NewMetadataProviderBuilder(config.Integrator)
	if err != nil {
		return nil, err
	}
	metadataProviderBuilder.ShoppingCartExtension = config.ShoppingCartExtension

	metadataProvider, err := metadataProviderBuilder.Build()
	if err != nil {
		return nil, err
	}

	authenticator, err := getAuthenticator(config)
	if err != nil {
		return nil, err
	}

	marshaller := json.DefaultMarshaller()

	return builder.
		WithAPIEndpoint(&config.APIEndpoint).
		WithConnection(connection).
		WithMetadataProvider(metadataProvider).
		WithAuthenticator(authenticator).
		WithMarshaller(marshaller), nil
}

func getAuthenticator(conf *configuration.CommunicatorConfiguration) (authentication.Authenticator, error) {
	if conf.AuthorizationType == configuration.OAUTH2 {
		return oauth2.NewAuthenticator(conf)
	}
	return nil, errors.New("unknown authorizationType " + string(conf.AuthorizationType))
}

// CreateCommunicatorFromConfiguration creates a Communicator with the given CommunicatorConfiguration
func CreateCommunicatorFromConfiguration(conf *configuration.CommunicatorConfiguration) (*communicator.Communicator, error) {
	builder, err := CreateCommunicatorBuilderFromConfiguration(conf)
	if err != nil {
		return nil, err
	}

	return builder.Build()
}

// CreateCommunicatorWithDefaultMarshaller creates a Communicator with the given components and a default marshaller
func CreateCommunicatorWithDefaultMarshaller(apiEndpoint *url.URL, connection communicator.Connection, authenticator authentication.Authenticator, metadataProvider *communicator.MetadataProvider) (*communicator.Communicator, error) {
	return communicator.NewCommunicator(apiEndpoint, connection, authenticator, metadataProvider, json.DefaultMarshaller())
}

// CreateClientFromConfiguration creates a Client with the given CommunicatorConfiguration
func CreateClientFromConfiguration(config *configuration.CommunicatorConfiguration) (*Client, error) {
	comm, err := CreateCommunicatorFromConfiguration(config)
	if err != nil {
		return nil, err
	}

	return CreateClientFromCommunicator(comm)
}

// CreateClientWithDefaultMarshaller creates a Client with the given components and a default marshaller
func CreateClientWithDefaultMarshaller(apiEndpoint *url.URL, connection communicator.Connection, authenticator authentication.Authenticator, metadataProvider *communicator.MetadataProvider) (*Client, error) {
	comm, err := CreateCommunicatorWithDefaultMarshaller(apiEndpoint, connection, authenticator, metadataProvider)
	if err != nil {
		return nil, err
	}

	return CreateClientFromCommunicator(comm)
}

// CreateClientFromCommunicator creates a Client with the given Communicator
func CreateClientFromCommunicator(communicator *communicator.Communicator) (*Client, error) {
	return NewClient(communicator)
}

// NewCallContext creates an empty CallContext
func NewCallContext() *communicator.CallContext {
	return communicator.NewCallContext()
}
