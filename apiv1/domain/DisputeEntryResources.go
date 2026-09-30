// This file was automatically generated.

package domain

// DisputeEntryResources represents class DisputeEntryResources.
type DisputeEntryResources struct {
	DisputeEntries *[]DisputeEntryWithDisputeSummary `json:"disputeEntries,omitempty"`
	Pagination     *PaginationResponse               `json:"pagination,omitempty"`
	RequestID      *string                           `json:"requestId,omitempty"`
}

// NewDisputeEntryResources constructs a new DisputeEntryResources instance.
func NewDisputeEntryResources() *DisputeEntryResources {
	return &DisputeEntryResources{}
}
