// This file was automatically generated.

package errors

import "strconv"

// AuthorizationError represents an error response from the Worldline Acquiring platform when API authorization failed.
type AuthorizationError struct {
	errorMessage string
	statusCode   int
	responseBody string
	typeValue    string
	title        string
	status       *int32
	detail       string
	instance     string
}

// Message returns the error message
func (e AuthorizationError) Message() string {
	return e.errorMessage
}

// StatusCode returns the status code
func (e AuthorizationError) StatusCode() int {
	return e.statusCode
}

// ResponseBody returns the response body
func (e AuthorizationError) ResponseBody() string {
	return e.responseBody
}

// Type implements the APIError interface
func (e AuthorizationError) Type() string {
	return e.typeValue
}

// Title implements the APIError interface
func (e AuthorizationError) Title() string {
	return e.title
}

// Status implements the APIError interface
func (e AuthorizationError) Status() *int32 {
	return e.status
}

// Detail implements the APIError interface
func (e AuthorizationError) Detail() string {
	return e.detail
}

// Instance implements the APIError interface
func (e AuthorizationError) Instance() string {
	return e.instance
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e AuthorizationError) String() string {
	list := e.errorMessage

	if e.statusCode > 0 {
		list = list + "; statusCode=" + strconv.Itoa(e.statusCode)
	}
	if len(e.responseBody) != 0 {
		list = list + "; responseBody='" + e.responseBody + "'"
	}

	return list
}

// Error implements the error interface
func (e AuthorizationError) Error() string {
	return e.String()
}

// NewAuthorizationError creates a new AuthorizationError with the given statusCode, responseBody and response fields
func NewAuthorizationError(statusCode int, responseBody, typeValue string, title string, status *int32, detail string, instance string) (*AuthorizationError, error) {
	return &AuthorizationError{"The Worldline Acquiring platform returned an API authorization error response", statusCode, responseBody, typeValue, title, status, detail, instance}, nil
}

// NewAuthorizationErrorVerbose creates a new AuthorizationError with the given message, statusCode and response fields
func NewAuthorizationErrorVerbose(message string, statusCode int, responseBody, typeValue string, title string, status *int32, detail string, instance string) (*AuthorizationError, error) {
	return &AuthorizationError{message, statusCode, responseBody, typeValue, title, status, detail, instance}, nil
}
