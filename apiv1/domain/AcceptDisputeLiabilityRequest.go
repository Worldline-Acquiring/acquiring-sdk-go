// This file was automatically generated.

package domain

// AcceptDisputeLiabilityRequest represents class AcceptDisputeLiabilityRequest.
type AcceptDisputeLiabilityRequest struct {
	IncludeEntries *bool   `json:"includeEntries,omitempty"`
	MessageText    *string `json:"messageText,omitempty"`
	UserID         *string `json:"userId,omitempty"`
}

// NewAcceptDisputeLiabilityRequest constructs a new AcceptDisputeLiabilityRequest instance.
func NewAcceptDisputeLiabilityRequest() *AcceptDisputeLiabilityRequest {
	return &AcceptDisputeLiabilityRequest{}
}
