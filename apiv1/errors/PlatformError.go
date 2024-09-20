// This file was automatically generated.

package errors

import "strconv"

// PlatformError represents an error response from the Worldline Acquiring platform when something went wrong at the Worldline Acquiring platform or further downstream.
type PlatformError struct {
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
func (e PlatformError) Message() string {
	return e.errorMessage
}

// StatusCode returns the status code
func (e PlatformError) StatusCode() int {
	return e.statusCode
}

// ResponseBody returns the response body
func (e PlatformError) ResponseBody() string {
	return e.responseBody
}

// Type implements the APIError interface
func (e PlatformError) Type() string {
	return e.typeValue
}

// Title implements the APIError interface
func (e PlatformError) Title() string {
	return e.title
}

// Status implements the APIError interface
func (e PlatformError) Status() *int32 {
	return e.status
}

// Detail implements the APIError interface
func (e PlatformError) Detail() string {
	return e.detail
}

// Instance implements the APIError interface
func (e PlatformError) Instance() string {
	return e.instance
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e PlatformError) String() string {
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
func (e PlatformError) Error() string {
	return e.String()
}

// NewPlatformError creates a new PlatformError with the given statusCode, responseBody and response fields
func NewPlatformError(statusCode int, responseBody, typeValue string, title string, status *int32, detail string, instance string) (*PlatformError, error) {
	return &PlatformError{"The Worldline Acquiring platform returned an error response", statusCode, responseBody, typeValue, title, status, detail, instance}, nil
}

// NewPlatformErrorVerbose creates a new PlatformError with the given message, statusCode and response fields
func NewPlatformErrorVerbose(message string, statusCode int, responseBody, typeValue string, title string, status *int32, detail string, instance string) (*PlatformError, error) {
	return &PlatformError{message, statusCode, responseBody, typeValue, title, status, detail, instance}, nil
}
