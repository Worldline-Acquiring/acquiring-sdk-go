// This file was automatically generated.

package refunds

import (
	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/errors"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
	commErrors "github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/errors"
)

// Client represents a Refunds client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// ProcessStandaloneRefund represents the resource /processing/v1/{acquirerId}/{merchantId}/refunds - Create standalone refund
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Refunds/operation/processStandaloneRefund
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
func (c *Client) ProcessStandaloneRefund(body domain.APIRefundRequest, context *communicator.CallContext) (domain.APIRefundResponse, error) {
	var resultObject domain.APIRefundResponse

	uri, err := c.apiResource.InstantiateURIWithContext("/processing/v1/{acquirerId}/{merchantId}/refunds", nil)
	if err != nil {
		return resultObject, err
	}

	postErr := c.apiResource.Communicator().Post(uri, nil, nil, body, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*commErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &domain.APIPaymentErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return resultObject, err
			}

			err, createErr := v1Errors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, postErr
	}

	return resultObject, nil
}

// GetRefund represents the resource /processing/v1/{acquirerId}/{merchantId}/refunds/{refundId} - Retrieve refund
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Refunds/operation/getRefund
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
func (c *Client) GetRefund(refundID string, query GetRefundParams, context *communicator.CallContext) (domain.APIRefundResource, error) {
	var resultObject domain.APIRefundResource

	pathContext := map[string]string{
		"refundId": refundID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/processing/v1/{acquirerId}/{merchantId}/refunds/{refundId}", pathContext)
	if err != nil {
		return resultObject, err
	}

	getErr := c.apiResource.Communicator().Get(uri, nil, &query, context, &resultObject)
	if getErr != nil {
		responseError, isResponseError := getErr.(*commErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &domain.APIPaymentErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return resultObject, err
			}

			err, createErr := v1Errors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, getErr
	}

	return resultObject, nil
}

// CaptureRefund represents the resource /processing/v1/{acquirerId}/{merchantId}/refunds/{refundId}/captures - Capture refund
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Refunds/operation/captureRefund
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
func (c *Client) CaptureRefund(refundID string, body domain.APICaptureRequestForRefund, context *communicator.CallContext) (domain.APIActionResponseForRefund, error) {
	var resultObject domain.APIActionResponseForRefund

	pathContext := map[string]string{
		"refundId": refundID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/processing/v1/{acquirerId}/{merchantId}/refunds/{refundId}/captures", pathContext)
	if err != nil {
		return resultObject, err
	}

	postErr := c.apiResource.Communicator().Post(uri, nil, nil, body, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*commErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &domain.APIPaymentErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return resultObject, err
			}

			err, createErr := v1Errors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, postErr
	}

	return resultObject, nil
}

// ReverseRefundAuthorization represents the resource /processing/v1/{acquirerId}/{merchantId}/refunds/{refundId}/authorization-reversals - Reverse refund authorization
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Refunds/operation/reverseRefundAuthorization
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
func (c *Client) ReverseRefundAuthorization(refundID string, body domain.APIPaymentReversalRequest, context *communicator.CallContext) (domain.APIActionResponseForRefund, error) {
	var resultObject domain.APIActionResponseForRefund

	pathContext := map[string]string{
		"refundId": refundID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/processing/v1/{acquirerId}/{merchantId}/refunds/{refundId}/authorization-reversals", pathContext)
	if err != nil {
		return resultObject, err
	}

	postErr := c.apiResource.Communicator().Post(uri, nil, nil, body, context, &resultObject)
	if postErr != nil {
		responseError, isResponseError := postErr.(*commErrors.ResponseError)
		if isResponseError {
			var errorObject interface{}

			errorObject = &domain.APIPaymentErrorResponse{}
			err = c.apiResource.Communicator().Marshaller().Unmarshal(responseError.Body(), errorObject)
			if err != nil {
				return resultObject, err
			}

			err, createErr := v1Errors.CreateAPIError(responseError.StatusCode(), responseError.Body(), errorObject, context)
			if createErr != nil {
				return resultObject, createErr
			}

			return resultObject, err
		}

		return resultObject, postErr
	}

	return resultObject, nil
}

// NewClient constructs a new Refunds client
//
// parent is the communicator.APIResource on top of which we want to build the new Refunds client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
