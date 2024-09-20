// This file was automatically generated.

package errors

import "strconv"

// ReferenceError represents an error response from the Worldline Acquiring platform when a non-existing or removed object is trying to be accessed.
type ReferenceError struct {
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
func (e ReferenceError) Message() string {
	return e.errorMessage
}

// StatusCode returns the status code
func (e ReferenceError) StatusCode() int {
	return e.statusCode
}

// ResponseBody returns the response body
func (e ReferenceError) ResponseBody() string {
	return e.responseBody
}

// Type implements the APIError interface
func (e ReferenceError) Type() string {
	return e.typeValue
}

// Title implements the APIError interface
func (e ReferenceError) Title() string {
	return e.title
}

// Status implements the APIError interface
func (e ReferenceError) Status() *int32 {
	return e.status
}

// Detail implements the APIError interface
func (e ReferenceError) Detail() string {
	return e.detail
}

// Instance implements the APIError interface
func (e ReferenceError) Instance() string {
	return e.instance
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e ReferenceError) String() string {
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
func (e ReferenceError) Error() string {
	return e.String()
}

// NewReferenceError creates a new ReferenceError with the given statusCode, responseBody and response fields
func NewReferenceError(statusCode int, responseBody, typeValue string, title string, status *int32, detail string, instance string) (*ReferenceError, error) {
	return &ReferenceError{"The Worldline Acquiring platform returned a reference error response", statusCode, responseBody, typeValue, title, status, detail, instance}, nil
}

// NewReferenceErrorVerbose creates a new ReferenceError with the given message, statusCode and response fields
func NewReferenceErrorVerbose(message string, statusCode int, responseBody, typeValue string, title string, status *int32, detail string, instance string) (*ReferenceError, error) {
	return &ReferenceError{message, statusCode, responseBody, typeValue, title, status, detail, instance}, nil
}
