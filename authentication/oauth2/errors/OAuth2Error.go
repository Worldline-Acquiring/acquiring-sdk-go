package errors

// OAuth2Error represents an error regarding the authorization with the Worldline OAuth2 Authorization Server
type OAuth2Error struct {
	errorMessage string
}

// String implements the Stringer interface
func (oe *OAuth2Error) String() string {
	return oe.errorMessage
}

// Error implements the error interface
func (oe *OAuth2Error) Error() string {
	return oe.String()
}

// NewOAuth2Error creates an OAuth2Error with the given message
func NewOAuth2Error(message string) *OAuth2Error {
	return &OAuth2Error{message}
}
