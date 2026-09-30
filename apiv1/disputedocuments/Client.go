// This file was automatically generated.

package disputedocuments

import (
	"errors"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/domain"
	v1Errors "github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/errors"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
	commErrors "github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/errors"
)

// Client represents a DisputeDocuments client. Thread-safe.
type Client struct {
	apiResource *communicator.APIResource
}

// UploadDisputeDocument represents the resource /dispute-management/v1/documents - Upload Dispute Document.
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Dispute-Documents/operation/uploadDisputeDocument.
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
func (c *Client) UploadDisputeDocument(body UploadDisputeDocumentRequest, context *communicator.CallContext) (domain.UploadDocumentResponse, error) {
	var resultObject domain.UploadDocumentResponse

	uri, err := c.apiResource.InstantiateURIWithContext("/dispute-management/v1/documents", nil)
	if err != nil {
		return resultObject, err
	}

	postErr := c.apiResource.Communicator().Post(uri, nil, nil, &body, context, &resultObject)
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

// GetDisputeDocument represents the resource /dispute-management/v1/disputes/{disputeId}/documents/{documentId} - Retrieve Dispute Document.
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Dispute-Documents/operation/getDisputeDocument.
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
func (c *Client) GetDisputeDocument(disputeID string, documentID string, context *communicator.CallContext, bodyHandler communicator.BodyHandler) error {
	pathContext := map[string]string{
		"disputeId":  disputeID,
		"documentId": documentID,
	}

	uri, err := c.apiResource.InstantiateURIWithContext("/dispute-management/v1/disputes/{disputeId}/documents/{documentId}", pathContext)
	if err != nil {
		return err
	}

	getErr := c.apiResource.Communicator().GetWithHandler(uri, nil, nil, context, bodyHandler)
	if getErr != nil {
		var responseError *commErrors.ResponseError
		if errors.As(getErr, &responseError) {
			errorObject := &domain.APIPaymentErrorResponse{}
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

// NewClient constructs a new DisputeDocuments client.
//
// parent is the communicator.APIResource on top of which we want to build the new DisputeDocuments client.
func NewClient(parent *communicator.APIResource, pathContext map[string]string) (*Client, error) {
	apiResource, err := communicator.NewAPIResourceWithParent(parent, pathContext)
	if err != nil {
		return nil, err
	}

	return &Client{apiResource}, nil
}
