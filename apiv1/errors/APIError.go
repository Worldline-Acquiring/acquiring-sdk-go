// This file was automatically generated.

package errors

// APIError represents an error response from the Worldline Acquiring platform.
type APIError interface {
	// Error implements the error interface
	Error() string

	// Message gets the raw response body that was returned by the Worldline Acquiring platform.
	Message() string

	// StatusCode gets the HTTP status code that was returned by the Worldline Acquiring platform.
	StatusCode() int

	// ResponseBody gets the raw response body that was returned by the Worldline Acquiring platform.
	ResponseBody() string

	// Type gets the type received from the Worldline Acquiring platform if available.
	Type() string

	// Title gets the title received from the Worldline Acquiring platform if available.
	Title() string

	// Status gets the status received from the Worldline Acquiring platform if available.
	Status() *int32

	// Detail gets the detail received from the Worldline Acquiring platform if available.
	Detail() string

	// Instance gets the instance received from the Worldline Acquiring platform if available.
	Instance() string
}
