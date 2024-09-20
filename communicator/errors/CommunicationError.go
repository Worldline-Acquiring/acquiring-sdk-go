package errors

// CommunicationError represents an error during the communication with the Worldline Acquiring platform
type CommunicationError struct {
	internalError error
}

// InternalError returns the internal error encountered
func (ce CommunicationError) InternalError() error {
	return ce.internalError
}

// Error implements the error interface
func (ce CommunicationError) Error() string {
	return ce.String()
}

// String implements the Stringer interface
// Format: 'There was an error in the communication with the Worldline Acquiring platform error'
func (ce CommunicationError) String() string {
	return "There was an error in the communication with the Worldline Acquiring platform " + ce.internalError.Error()
}

// NewCommunicationError creates a CommunicationError with the given internal error
func NewCommunicationError(internalError error) (*CommunicationError, error) {
	return &CommunicationError{internalError}, nil
}
