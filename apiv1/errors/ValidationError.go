// This file was automatically generated.

package errors

import "strconv"

// ValidationError represents an error response from the Worldline Acquiring platform when validation of requests failed.
type ValidationError struct {
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
func (e ValidationError) Message() string {
	return e.errorMessage
}

// StatusCode returns the status code
func (e ValidationError) StatusCode() int {
	return e.statusCode
}

// ResponseBody returns the response body
func (e ValidationError) ResponseBody() string {
	return e.responseBody
}

// Type implements the APIError interface
func (e ValidationError) Type() string {
	return e.typeValue
}

// Title implements the APIError interface
func (e ValidationError) Title() string {
	return e.title
}

// Status implements the APIError interface
func (e ValidationError) Status() *int32 {
	return e.status
}

// Detail implements the APIError interface
func (e ValidationError) Detail() string {
	return e.detail
}

// Instance implements the APIError interface
func (e ValidationError) Instance() string {
	return e.instance
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e ValidationError) String() string {
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
func (e ValidationError) Error() string {
	return e.String()
}

// NewValidationError creates a new ValidationError with the given statusCode, responseBody and response fields
func NewValidationError(statusCode int, responseBody, typeValue string, title string, status *int32, detail string, instance string) (*ValidationError, error) {
	return &ValidationError{"The Worldline Acquiring platform returned an incorrect request error response", statusCode, responseBody, typeValue, title, status, detail, instance}, nil
}

// NewValidationErrorVerbose creates a new ValidationError with the given message, statusCode and response fields
func NewValidationErrorVerbose(message string, statusCode int, responseBody, typeValue string, title string, status *int32, detail string, instance string) (*ValidationError, error) {
	return &ValidationError{message, statusCode, responseBody, typeValue, title, status, detail, instance}, nil
}
