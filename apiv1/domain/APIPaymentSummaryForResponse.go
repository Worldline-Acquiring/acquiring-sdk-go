// This file was automatically generated.

package domain

// APIPaymentSummaryForResponse represents class ApiPaymentSummaryForResponse
type APIPaymentSummaryForResponse struct {
	PaymentID       *string                    `json:"paymentId,omitempty"`
	References      *APIReferencesForResponses `json:"references,omitempty"`
	RetryAfter      *string                    `json:"retryAfter,omitempty"`
	Status          *string                    `json:"status,omitempty"`
	StatusTimestamp *string                    `json:"statusTimestamp,omitempty"`
}

// NewAPIPaymentSummaryForResponse constructs a new APIPaymentSummaryForResponse instance
func NewAPIPaymentSummaryForResponse() *APIPaymentSummaryForResponse {
	return &APIPaymentSummaryForResponse{}
}
