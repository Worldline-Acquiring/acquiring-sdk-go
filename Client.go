// This file was automatically generated.

package acquiringsdk

import (
	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/logging"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/logging/obfuscation"
)

// Client is the Worldline Acquiring platform client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// SetBodyObfuscator sets the body obfuscator to use.
func (c *Client) SetBodyObfuscator(bodyObfuscator obfuscation.BodyObfuscator) {
	c.apiResource.Communicator().SetBodyObfuscator(bodyObfuscator)
}

// SetHeaderObfuscator sets the header obfuscator to use.
func (c *Client) SetHeaderObfuscator(headerObfuscator obfuscation.HeaderObfuscator) {
	c.apiResource.Communicator().SetHeaderObfuscator(headerObfuscator)
}

// EnableLogging turns on logging using the given communicator logger.
func (c *Client) EnableLogging(communicatorLogger logging.CommunicatorLogger) {
	c.apiResource.Communicator().EnableLogging(communicatorLogger)
}

// DisableLogging turns off logging.
func (c *Client) DisableLogging() {
	c.apiResource.Communicator().DisableLogging()
}

// Close calls the internal closer of the communicator
func (c *Client) Close() error {
	return c.apiResource.Communicator().Close()
}

// V1 represents API v1
func (c *Client) V1() *apiv1.Client {
	client, _ := apiv1.NewClient(c.apiResource, nil)
	return client
}

// NewClient creates a new Client with the given communicator
func NewClient(comm *communicator.Communicator) (*Client, error) {
	apiResource, err := communicator.NewAPIResource(comm, nil)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
