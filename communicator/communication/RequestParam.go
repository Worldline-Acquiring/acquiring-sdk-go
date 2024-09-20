package communication

import "errors"

// RequestParam represents a single request parameter. Immutable.
type RequestParam struct {
	name, value string
}

// RequestParams represents a slice of RequestParam
type RequestParams []RequestParam

// Name returns the name of the RequestParam
func (rp RequestParam) Name() string {
	return rp.name
}

// Value returns the value of the RequestParam
func (rp RequestParam) Value() string {
	return rp.value
}

// String is the implementation of the Stringer interface
// Format: 'name:value'
func (rp RequestParam) String() string {
	return rp.name + ":" + rp.value
}

// NewRequestParam creates a RequestParam with the given name and value
func NewRequestParam(name, value string) (*RequestParam, error) {
	if len(name) == 0 {
		return nil, errors.New("name is required")
	}

	return &RequestParam{name, value}, nil
}
