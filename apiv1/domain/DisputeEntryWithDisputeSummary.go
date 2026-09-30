// This file was automatically generated.

package domain

// DisputeEntryWithDisputeSummary represents class DisputeEntryWithDisputeSummary.
type DisputeEntryWithDisputeSummary struct {
	DisputeEntry
	DisputeID      *string         `json:"disputeId,omitempty"`
	DisputeSummary *DisputeSummary `json:"disputeSummary,omitempty"`
}

// NewDisputeEntryWithDisputeSummary constructs a new DisputeEntryWithDisputeSummary instance.
func NewDisputeEntryWithDisputeSummary() *DisputeEntryWithDisputeSummary {
	return &DisputeEntryWithDisputeSummary{}
}
