// This file was automatically generated.

package ping

import (
	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/errors"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
	commErrors "github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/errors"
)

// Client represents a Ping client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// Ping represents the resource /services/v1/ping - Check API connection
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Ping/operation/ping
//
// Can return any of the following errors:
//   * ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
//   * AuthorizationError if the request was not allowed (HTTP status code 403)
//   * ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//     or there was a conflict (HTTP status code 404, 409 or 410)
//   * PlatformError if something went wrong at the Worldline Acquiring platform,
//     the Worldline Acquiring platform was unable to process a message from a downstream partner/acquirer,
//     or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
//   * APIError if the Worldline Acquiring platform returned any other error
func (c *Client) Ping(context *communicator.CallContext) error {
	uri, err := c.apiResource.InstantiateURIWithContext("/services/v1/ping", nil)
	if err != nil {
		return err
	}

	var resultObject map[string]interface{}
	getErr := c.apiResource.Communicator().Get(uri, nil, nil, context, &resultObject)
	if getErr != nil {
		responseError, isResponseError := getErr.(*commErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &domain.APIPaymentErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return err
			}

			err, createErr := v1Errors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return createErr
			}

			return err
		}

		return getErr
	}

	return nil
}

// NewClient constructs a new Ping client
//
// parent is the communicator.APIResource on top of which we want to build the new Ping client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
