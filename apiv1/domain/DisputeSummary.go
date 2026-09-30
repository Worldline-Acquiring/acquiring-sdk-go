// This file was automatically generated.

package domain

// DisputeSummary represents class DisputeSummary.
type DisputeSummary struct {
	AcquirerDisputeReference *string                         `json:"acquirerDisputeReference,omitempty"`
	MerchantData             *DisputeMerchantDataBase        `json:"merchantData,omitempty"`
	OriginalTransactionData  *OriginalTransactionSummaryData `json:"originalTransactionData,omitempty"`
	SchemeReason             *string                         `json:"schemeReason,omitempty"`
	SchemeReasonDescription  *string                         `json:"schemeReasonDescription,omitempty"`
	UnifiedCategory          *string                         `json:"unifiedCategory,omitempty"`
	UnifiedReason            *string                         `json:"unifiedReason,omitempty"`
}

// NewDisputeSummary constructs a new DisputeSummary instance.
func NewDisputeSummary() *DisputeSummary {
	return &DisputeSummary{}
}
