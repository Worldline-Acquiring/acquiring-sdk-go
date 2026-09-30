// This file was automatically generated.

package domain

// DisputeReferences represents class DisputeReferences.
type DisputeReferences struct {
	AcquirerDisputeReference *string `json:"acquirerDisputeReference,omitempty"`
	SchemeDisputeReference   *string `json:"schemeDisputeReference,omitempty"`
}

// NewDisputeReferences constructs a new DisputeReferences instance.
func NewDisputeReferences() *DisputeReferences {
	return &DisputeReferences{}
}
