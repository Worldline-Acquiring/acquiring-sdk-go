// This file was automatically generated.

package domain

// DccProposal represents class DccProposal
type DccProposal struct {
	OriginalAmount  *AmountData `json:"originalAmount,omitempty"`
	Rate            *RateData   `json:"rate,omitempty"`
	RateReferenceID *string     `json:"rateReferenceId,omitempty"`
	ResultingAmount *AmountData `json:"resultingAmount,omitempty"`
}

// NewDccProposal constructs a new DccProposal instance
func NewDccProposal() *DccProposal {
	return &DccProposal{}
}
