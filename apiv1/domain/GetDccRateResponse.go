// This file was automatically generated.

package domain

// GetDccRateResponse represents class GetDccRateResponse
type GetDccRateResponse struct {
	DisclaimerDisplay *string      `json:"disclaimerDisplay,omitempty"`
	DisclaimerReceipt *string      `json:"disclaimerReceipt,omitempty"`
	Proposal          *DccProposal `json:"proposal,omitempty"`
	Result            *string      `json:"result,omitempty"`
}

// NewGetDccRateResponse constructs a new GetDccRateResponse instance
func NewGetDccRateResponse() *GetDccRateResponse {
	return &GetDccRateResponse{}
}
