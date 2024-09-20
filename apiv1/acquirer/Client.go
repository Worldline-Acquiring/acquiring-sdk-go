// This file was automatically generated.

package acquirer

import (
	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/acquirer/merchant"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
)

// Client represents a Acquirer client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// Merchant represents the resource /processing/v1/{acquirerId}/{merchantId}
func (c *Client) Merchant(merchantID string) *merchant.Client {
	client, _ := merchant.NewClient(c.apiResource, map[string]string{
		"merchantId": merchantID,
	})
	return client
}

// NewClient constructs a new Acquirer client
//
// parent is the communicator.APIResource on top of which we want to build the new Acquirer client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
