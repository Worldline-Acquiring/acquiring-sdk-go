package communicator

import (
	"io"
	"net/url"
	"time"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/logging"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/logging/obfuscation"
)

// Connection represents a pooled connection to the Worldline Acquiring platform server.
// Instead of setting up a new HTTP connection for each request, this connection uses a pool of HTTP connections.
// Thread-safe
type Connection interface {
	logging.Capable
	obfuscation.Capable
	io.Closer

	// Get sends a GET request to the Worldline Acquiring platform and calls the given response handler with the response.
	Get(resourceURI url.URL, requestHeaders []communication.Header, handler communication.ResponseHandler) (interface{}, error)

	// Delete sends a DELETE request to the Worldline Acquiring platform and calls the given response handler with the response.
	Delete(resourceURI url.URL, requestHeaders []communication.Header, handler communication.ResponseHandler) (interface{}, error)

	// Post sends a POST request to the Worldline Acquiring platform and calls the given response handler with the response.
	Post(resourceURI url.URL, requestHeaders []communication.Header, body string, handler communication.ResponseHandler) (interface{}, error)

	// PostMultipart sends a multipart/form-data POST request to the Worldline Acquiring platform and calls the given response handler with the response.
	PostMultipart(resourceURI url.URL, requestHeaders []communication.Header, body *communication.MultipartFormDataObject, handler communication.ResponseHandler) (interface{}, error)

	// Put sends a PUT request to the Worldline Acquiring platform and calls the given response handler with the response.
	Put(resourceURI url.URL, requestHeaders []communication.Header, body string, handler communication.ResponseHandler) (interface{}, error)

	// PutMultipart sends a multipart/form-data PUT request to the Worldline Acquiring platform and calls the given response handler with the response.
	PutMultipart(resourceURI url.URL, requestHeaders []communication.Header, body *communication.MultipartFormDataObject, hHandler communication.ResponseHandler) (interface{}, error)

	// CloseIdleConnections closes all HTTP connections that have been idle for the specified time.
	// This should also include all expired HTTP connections.
	// timespan represents the time spent idle
	CloseIdleConnections(time time.Duration)

	// CloseExpiredConnections closes all expired HTTP connections.
	CloseExpiredConnections()
}
