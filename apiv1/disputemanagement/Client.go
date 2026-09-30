// This file was automatically generated.

package disputemanagement

import (
	"errors"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/errors"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
	commErrors "github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/errors"
)

// Client represents a DisputeManagement client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// SearchDisputes represents the resource /dispute-management/v1/disputes/search - Search Disputes.
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Dispute-Management/operation/searchDisputes.
//
// Can return any of the following errors:
//   - ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
//   - AuthorizationError if the request was not allowed (HTTP status code 403)
//   - ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//     or there was a conflict (HTTP status code 404, 409 or 410)
//   - PlatformError if something went wrong at the Worldline Acquiring platform,
//     the Worldline Acquiring platform was unable to process a message from a downstream partner/acquirer,
//     or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
//   - APIError if the Worldline Acquiring platform returned any other error
func (c *Client) SearchDisputes(body domain.SearchDisputesRequest, context *communicator.CallContext) (domain.SearchDisputesResponse, error) {
	var resultObject domain.SearchDisputesResponse

	uri, err := c.apiResource.InstantiateURIWithContext("/dispute-management/v1/disputes/search", nil)
	if err != nil {
		return resultObject, err
	}

	postErr := c.apiResource.Communicator().Post(uri, nil, nil, body, context, &resultObject)
	if postErr != nil {
		var responseError *commErrors.ResponseError
		if errors.As(postErr, &responseError) {
			errorObject := &domain.APIPaymentErrorResponse{}
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

// GetDispute represents the resource /dispute-management/v1/disputes/{disputeId} - Retrieve Dispute.
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Dispute-Management/operation/getDispute.
//
// Can return any of the following errors:
//   - ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
//   - AuthorizationError if the request was not allowed (HTTP status code 403)
//   - ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//     or there was a conflict (HTTP status code 404, 409 or 410)
//   - PlatformError if something went wrong at the Worldline Acquiring platform,
//     the Worldline Acquiring platform was unable to process a message from a downstream partner/acquirer,
//     or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
//   - APIError if the Worldline Acquiring platform returned any other error
func (c *Client) GetDispute(disputeID string, query GetDisputeParams, context *communicator.CallContext) (domain.DisputeResponse, error) {
	var resultObject domain.DisputeResponse

	pathContext := map[string]string{
		"disputeId": disputeID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/dispute-management/v1/disputes/{disputeId}", pathContext)
	if err != nil {
		return resultObject, err
	}

	getErr := c.apiResource.Communicator().Get(uri, nil, &query, context, &resultObject)
	if getErr != nil {
		var responseError *commErrors.ResponseError
		if errors.As(getErr, &responseError) {
			errorObject := &domain.APIPaymentErrorResponse{}
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

// AcceptDisputeLiability represents the resource /dispute-management/v1/disputes/{disputeId}/accept - Accept Liability.
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Dispute-Management/operation/acceptDisputeLiability.
//
// Can return any of the following errors:
//   - ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
//   - AuthorizationError if the request was not allowed (HTTP status code 403)
//   - ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//     or there was a conflict (HTTP status code 404, 409 or 410)
//   - PlatformError if something went wrong at the Worldline Acquiring platform,
//     the Worldline Acquiring platform was unable to process a message from a downstream partner/acquirer,
//     or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
//   - APIError if the Worldline Acquiring platform returned any other error
func (c *Client) AcceptDisputeLiability(disputeID string, body domain.AcceptDisputeLiabilityRequest, context *communicator.CallContext) (domain.DisputeResponse, error) {
	var resultObject domain.DisputeResponse

	pathContext := map[string]string{
		"disputeId": disputeID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/dispute-management/v1/disputes/{disputeId}/accept", pathContext)
	if err != nil {
		return resultObject, err
	}

	postErr := c.apiResource.Communicator().Post(uri, nil, nil, body, context, &resultObject)
	if postErr != nil {
		var responseError *commErrors.ResponseError
		if errors.As(postErr, &responseError) {
			errorObject := &domain.APIPaymentErrorResponse{}
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

// SubmitEvidence represents the resource /dispute-management/v1/disputes/{disputeId}/submit-evidence - Submit Evidence.
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Dispute-Management/operation/submitEvidence.
//
// Can return any of the following errors:
//   - ValidationError if the request was not correct and couldn't be processed (HTTP status code 400)
//   - AuthorizationError if the request was not allowed (HTTP status code 403)
//   - ReferenceError if an object was attempted to be referenced that doesn't exist or has been removed,
//     or there was a conflict (HTTP status code 404, 409 or 410)
//   - PlatformError if something went wrong at the Worldline Acquiring platform,
//     the Worldline Acquiring platform was unable to process a message from a downstream partner/acquirer,
//     or the service that you're trying to reach is temporary unavailable (HTTP status code 500, 502 or 503)
//   - APIError if the Worldline Acquiring platform returned any other error
func (c *Client) SubmitEvidence(disputeID string, body domain.SubmitEvidenceRequest, context *communicator.CallContext) (domain.DisputeResponse, error) {
	var resultObject domain.DisputeResponse

	pathContext := map[string]string{
		"disputeId": disputeID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/dispute-management/v1/disputes/{disputeId}/submit-evidence", pathContext)
	if err != nil {
		return resultObject, err
	}

	postErr := c.apiResource.Communicator().Post(uri, nil, nil, body, context, &resultObject)
	if postErr != nil {
		var responseError *commErrors.ResponseError
		if errors.As(postErr, &responseError) {
			errorObject := &domain.APIPaymentErrorResponse{}
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

// NewClient constructs a new DisputeManagement client.
//
// parent is the communicator.APIResource on top of which we want to build the new DisputeManagement client.
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
