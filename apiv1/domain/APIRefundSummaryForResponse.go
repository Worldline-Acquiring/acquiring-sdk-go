// This file was automatically generated.

package domain

// APIRefundSummaryForResponse represents class ApiRefundSummaryForResponse
type APIRefundSummaryForResponse struct {
	References      *APIReferencesForResponses `json:"references,omitempty"`
	RefundID        *string                    `json:"refundId,omitempty"`
	RetryAfter      *string                    `json:"retryAfter,omitempty"`
	Status          *string                    `json:"status,omitempty"`
	StatusTimestamp *string                    `json:"statusTimestamp,omitempty"`
}

// NewAPIRefundSummaryForResponse constructs a new APIRefundSummaryForResponse instance
func NewAPIRefundSummaryForResponse() *APIRefundSummaryForResponse {
	return &APIRefundSummaryForResponse{}
}
