// This file was automatically generated.

package domain

// DisputeCaseWithEntries represents class DisputeCaseWithEntries.
type DisputeCaseWithEntries struct {
	DisputeCase
	Entries *[]DisputeEntry `json:"entries,omitempty"`
}

// NewDisputeCaseWithEntries constructs a new DisputeCaseWithEntries instance.
func NewDisputeCaseWithEntries() *DisputeCaseWithEntries {
	return &DisputeCaseWithEntries{}
}
