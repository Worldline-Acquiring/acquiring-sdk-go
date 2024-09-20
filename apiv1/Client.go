// This file was automatically generated.

package apiv1

import (
	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/acquirer"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/ping"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
)

// Client represents a v1 client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// Acquirer represents the resource /processing/v1/{acquirerId}
func (c *Client) Acquirer(acquirerID string) *acquirer.Client {
	client, _ := acquirer.NewClient(c.apiResource, map[string]string{
		"acquirerId": acquirerID,
	})
	return client
}

// Ping represents the resource /services/v1/ping
func (c *Client) Ping() *ping.Client {
	client, _ := ping.NewClient(c.apiResource, nil)
	return client
}

// NewClient constructs a new v1 client
//
// parent is the communicator.APIResource on top of which we want to build the new v1 client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
