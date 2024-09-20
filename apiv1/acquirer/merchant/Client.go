// This file was automatically generated.

package merchant

import (
	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/acquirer/merchant/accountverifications"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/acquirer/merchant/dynamiccurrencyconversion"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/acquirer/merchant/payments"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/acquirer/merchant/refunds"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/acquirer/merchant/technicalreversals"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
)

// Client represents a Merchant client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// Payments represents the resource /processing/v1/{acquirerId}/{merchantId}/payments
func (c *Client) Payments() *payments.Client {
	client, _ := payments.NewClient(c.apiResource, nil)
	return client
}

// Refunds represents the resource /processing/v1/{acquirerId}/{merchantId}/refunds
func (c *Client) Refunds() *refunds.Client {
	client, _ := refunds.NewClient(c.apiResource, nil)
	return client
}

// AccountVerifications represents the resource /processing/v1/{acquirerId}/{merchantId}/account-verifications
func (c *Client) AccountVerifications() *accountverifications.Client {
	client, _ := accountverifications.NewClient(c.apiResource, nil)
	return client
}

// TechnicalReversals represents the resource /processing/v1/{acquirerId}/{merchantId}/operations/{operationId}/reverse
func (c *Client) TechnicalReversals() *technicalreversals.Client {
	client, _ := technicalreversals.NewClient(c.apiResource, nil)
	return client
}

// DynamicCurrencyConversion represents the resource /services/v1/{acquirerId}/{merchantId}/dcc-rates
func (c *Client) DynamicCurrencyConversion() *dynamiccurrencyconversion.Client {
	client, _ := dynamiccurrencyconversion.NewClient(c.apiResource, nil)
	return client
}

// NewClient constructs a new Merchant client
//
// parent is the communicator.APIResource on top of which we want to build the new Merchant client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
