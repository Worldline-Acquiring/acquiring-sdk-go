// This file was automatically generated.

package payments

import (
	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/errors"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
	commErrors "github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/errors"
)

// Client represents a Payments client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// ProcessPayment represents the resource /processing/v1/{acquirerId}/{merchantId}/payments - Create payment
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Payments/operation/processPayment
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
func (c *Client) ProcessPayment(body domain.APIPaymentRequest, context *communicator.CallContext) (domain.APIPaymentResponse, error) {
	var resultObject domain.APIPaymentResponse

	uri, err := c.apiResource.InstantiateURIWithContext("/processing/v1/{acquirerId}/{merchantId}/payments", nil)
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

// GetPaymentStatus represents the resource /processing/v1/{acquirerId}/{merchantId}/payments/{paymentId} - Retrieve payment
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Payments/operation/getPaymentStatus
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
func (c *Client) GetPaymentStatus(paymentID string, query GetPaymentStatusParams, context *communicator.CallContext) (domain.APIPaymentResource, error) {
	var resultObject domain.APIPaymentResource

	pathContext := map[string]string{
		"paymentId": paymentID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/processing/v1/{acquirerId}/{merchantId}/payments/{paymentId}", pathContext)
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

// SimpleCaptureOfPayment represents the resource /processing/v1/{acquirerId}/{merchantId}/payments/{paymentId}/captures - Capture payment
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Payments/operation/simpleCaptureOfPayment
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
func (c *Client) SimpleCaptureOfPayment(paymentID string, body domain.APICaptureRequest, context *communicator.CallContext) (domain.APIActionResponse, error) {
	var resultObject domain.APIActionResponse

	pathContext := map[string]string{
		"paymentId": paymentID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/processing/v1/{acquirerId}/{merchantId}/payments/{paymentId}/captures", pathContext)
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

// ReverseAuthorization represents the resource /processing/v1/{acquirerId}/{merchantId}/payments/{paymentId}/authorization-reversals - Reverse authorization
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Payments/operation/reverseAuthorization
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
func (c *Client) ReverseAuthorization(paymentID string, body domain.APIPaymentReversalRequest, context *communicator.CallContext) (domain.APIReversalResponse, error) {
	var resultObject domain.APIReversalResponse

	pathContext := map[string]string{
		"paymentId": paymentID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/processing/v1/{acquirerId}/{merchantId}/payments/{paymentId}/authorization-reversals", pathContext)
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

// IncrementPayment represents the resource /processing/v1/{acquirerId}/{merchantId}/payments/{paymentId}/increments - Increment authorization
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Payments/operation/incrementPayment
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
func (c *Client) IncrementPayment(paymentID string, body domain.APIIncrementRequest, context *communicator.CallContext) (domain.APIIncrementResponse, error) {
	var resultObject domain.APIIncrementResponse

	pathContext := map[string]string{
		"paymentId": paymentID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/processing/v1/{acquirerId}/{merchantId}/payments/{paymentId}/increments", pathContext)
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

// CreateRefund represents the resource /processing/v1/{acquirerId}/{merchantId}/payments/{paymentId}/refunds - Refund payment
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Payments/operation/createRefund
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
func (c *Client) CreateRefund(paymentID string, body domain.APIPaymentRefundRequest, context *communicator.CallContext) (domain.APIActionResponseForRefund, error) {
	var resultObject domain.APIActionResponseForRefund

	pathContext := map[string]string{
		"paymentId": paymentID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/processing/v1/{acquirerId}/{merchantId}/payments/{paymentId}/refunds", pathContext)
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

// NewClient constructs a new Payments client
//
// parent is the communicator.APIResource on top of which we want to build the new Payments client
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
