package authentication

import (
	"net/url"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
)

// Authenticator is the interface used to authenticate requests to the Worldline Acquiring platform. Thread-safe.
type Authenticator interface {

	// GetAuthorization returns a value that can be used for the "Authorization" header.
	// Note that the list of Request headers may not be modified and may not contain headers with the same name.
	GetAuthorization(httpMethod string, resourceURI url.URL, requestHeaders []communication.Header) (string, error)
}
