package acquiringsdk

import (
	"net/url"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/authentication"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/json"
)

// CommunicatorBuilder is the builder for Communicator objects
type CommunicatorBuilder struct {
	APIEndpoint      *url.URL
	Connection       communicator.Connection
	MetadataProvider *communicator.MetadataProvider
	Authenticator    authentication.Authenticator
	Marshaller       json.Marshaller
}

// WithAPIEndpoint sets the Worldline Acquiring platform API endpoint to be used by the Communicator
func (c *CommunicatorBuilder) WithAPIEndpoint(endpoint *url.URL) *CommunicatorBuilder {
	c.APIEndpoint = endpoint

	return c
}

// WithConnection sets the Connection to be used by the Communicator
func (c *CommunicatorBuilder) WithConnection(connection communicator.Connection) *CommunicatorBuilder {
	c.Connection = connection

	return c
}

// WithMetadataProvider sets the MetadataProvider to be used by the Communicator
func (c *CommunicatorBuilder) WithMetadataProvider(provider *communicator.MetadataProvider) *CommunicatorBuilder {
	c.MetadataProvider = provider

	return c
}

// WithAuthenticator sets the Authenticator to be used by the Communicator
func (c *CommunicatorBuilder) WithAuthenticator(auth authentication.Authenticator) *CommunicatorBuilder {
	c.Authenticator = auth

	return c
}

// WithMarshaller sets the Marshaller to be used by the Communicator
func (c *CommunicatorBuilder) WithMarshaller(marshaller json.Marshaller) *CommunicatorBuilder {
	c.Marshaller = marshaller

	return c
}

// Build creates a Communicator object based on the builder parameters
func (c *CommunicatorBuilder) Build() (*communicator.Communicator, error) {
	return communicator.NewCommunicator(c.APIEndpoint, c.Connection, c.Authenticator, c.MetadataProvider, c.Marshaller)
}

// NewCommunicatorBuilder creates a CommunicatorBuilder object
func NewCommunicatorBuilder() *CommunicatorBuilder {
	return &CommunicatorBuilder{}
}
