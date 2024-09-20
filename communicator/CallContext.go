package communicator

// CallContext can be used to send extra information with a request, and to receive extra information from a response.
// Please note that this type is not thread-safe. Each request should get its own call context instance.
type CallContext struct {
}

// NewCallContext creates an empty CallContext
func NewCallContext() *CallContext {
	return &CallContext{}
}
