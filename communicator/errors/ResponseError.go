package errors

import (
	"strconv"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
)

// ResponseError is returned when a response was received from the Worldline Acquiring platform which indicates an error.
type ResponseError struct {
	statusCode int
	body       string
	headers    []communication.Header
}

// StatusCode gets the HTTP status code that was returned by the Worldline Acquiring platform.
func (e ResponseError) StatusCode() int {
	return e.statusCode
}

// Body gets the raw response body that was returned by the Worldline Acquiring platform.
func (e ResponseError) Body() string {
	return e.body
}

// Headers gets the headers that were returned by the Worldline Acquiring platform.
func (e ResponseError) Headers() []communication.Header {
	return e.headers
}

// GetHeader returns the header with the given name, or nil if there was no such header.
func (e ResponseError) GetHeader(headerName string) *communication.Header {
	return communication.Headers(e.headers).GetHeader(headerName)
}

// String implements the Stringer interface
// Format: 'errorMessage; statusCode=; responseBody='
func (e ResponseError) String() string {
	result := "the Worldline Acquiring platform returned an error response"

	statusCode := e.statusCode
	if statusCode > 0 {
		result += ";  statusCode=" + strconv.Itoa(e.statusCode)
	}

	responseBody := e.body
	if len(responseBody) > 0 {
		result += "; responseBody='" + e.body + "'"
	}

	return result
}

// Error implements the error interface
func (e ResponseError) Error() string {
	return e.String()
}

// NewResponseError creates a new ResponseError with the specified response
func NewResponseError(statusCode int, body string, headers []communication.Header) (*ResponseError, error) {
	return &ResponseError{statusCode: statusCode, body: body, headers: headers}, nil
}
