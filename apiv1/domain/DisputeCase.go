// This file was automatically generated.

package domain

// DisputeCase represents class DisputeCase.
type DisputeCase struct {
	DisputeDateTimeData     *DisputeDateTimeData     `json:"disputeDateTimeData,omitempty"`
	DisputeID               *string                  `json:"disputeId,omitempty"`
	DisputeReferences       *DisputeReferences       `json:"disputeReferences,omitempty"`
	DisputeStage            *string                  `json:"disputeStage,omitempty"`
	DisputeStatus           *string                  `json:"disputeStatus,omitempty"`
	DisputeStatusCategory   *string                  `json:"disputeStatusCategory,omitempty"`
	IsOpen                  *bool                    `json:"isOpen,omitempty"`
	MerchantBalanceAmount   *SignedAmountData        `json:"merchantBalanceAmount,omitempty"`
	MerchantData            *DisputeMerchantData     `json:"merchantData,omitempty"`
	OriginalDisputeAmount   *AmountData              `json:"originalDisputeAmount,omitempty"`
	OriginalTransactionData *OriginalTransactionData `json:"originalTransactionData,omitempty"`
	SchemeReason            *string                  `json:"schemeReason,omitempty"`
	SchemeReasonDescription *string                  `json:"schemeReasonDescription,omitempty"`
	UnifiedCategory         *string                  `json:"unifiedCategory,omitempty"`
	UnifiedReason           *string                  `json:"unifiedReason,omitempty"`
}

// NewDisputeCase constructs a new DisputeCase instance.
func NewDisputeCase() *DisputeCase {
	return &DisputeCase{}
}
