// This file was automatically generated.

package errors

import (
	"net/http"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/apiv1/domain"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator"
)

// CreateAPIError is used internally in order to create an API error after an HTTP request is done
func CreateAPIError(statusCode int, responseBody string, errorObject interface{}, context *communicator.CallContext) (APIError, error) {
	var typeValue string
	var title string
	var status *int32
	var detail string
	var instance string

	switch response := errorObject.(type) {
	case *domain.APIPaymentErrorResponse:
		{
			if response.Type != nil {
				typeValue = *response.Type
			}
			if response.Title != nil {
				title = *response.Title
			}
			if response.Status != nil {
				status = response.Status
			}
			if response.Detail != nil {
				detail = *response.Detail
			}
			if response.Instance != nil {
				instance = *response.Instance
			}

			break
		}
	}

	switch statusCode {
	case http.StatusBadRequest:
		{
			return NewValidationError(statusCode, responseBody, typeValue, title, status, detail, instance)
		}
	case http.StatusForbidden:
		{
			return NewAuthorizationError(statusCode, responseBody, typeValue, title, status, detail, instance)
		}
	case http.StatusNotFound:
		{
			return NewReferenceError(statusCode, responseBody, typeValue, title, status, detail, instance)
		}
	case http.StatusConflict:
		{
			return NewReferenceError(statusCode, responseBody, typeValue, title, status, detail, instance)
		}
	case http.StatusGone:
		{
			return NewReferenceError(statusCode, responseBody, typeValue, title, status, detail, instance)
		}
	case http.StatusInternalServerError:
		{
			return NewPlatformError(statusCode, responseBody, typeValue, title, status, detail, instance)
		}
	case http.StatusBadGateway:
		{
			return NewPlatformError(statusCode, responseBody, typeValue, title, status, detail, instance)
		}
	case http.StatusServiceUnavailable:
		{
			return NewPlatformError(statusCode, responseBody, typeValue, title, status, detail, instance)
		}
	default:
		{
			return NewPlatformError(statusCode, responseBody, typeValue, title, status, detail, instance)
		}
	}
}
