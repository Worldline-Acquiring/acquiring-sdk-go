// This file was automatically generated.

package domain

// DisputeResponse represents class DisputeResponse.
type DisputeResponse struct {
	Dispute   *DisputeCaseWithEntries `json:"dispute,omitempty"`
	RequestID *string                 `json:"requestId,omitempty"`
}

// NewDisputeResponse constructs a new DisputeResponse instance.
func NewDisputeResponse() *DisputeResponse {
	return &DisputeResponse{}
}
