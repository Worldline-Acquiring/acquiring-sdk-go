// This file was automatically generated.

package domain

// APIRefundResource represents class ApiRefundResource
type APIRefundResource struct {
	CardPaymentData          *CardPaymentDataForResource `json:"cardPaymentData,omitempty"`
	InitialAuthorizationCode *string                     `json:"initialAuthorizationCode,omitempty"`
	Operations               *[]SubOperationForRefund    `json:"operations,omitempty"`
	ReferencedPaymentID      *string                     `json:"referencedPaymentId,omitempty"`
	References               *APIReferencesForResponses  `json:"references,omitempty"`
	RefundID                 *string                     `json:"refundId,omitempty"`
	RetryAfter               *string                     `json:"retryAfter,omitempty"`
	Status                   *string                     `json:"status,omitempty"`
	StatusTimestamp          *string                     `json:"statusTimestamp,omitempty"`
	TotalAuthorizedAmount    *AmountData                 `json:"totalAuthorizedAmount,omitempty"`
}

// NewAPIRefundResource constructs a new APIRefundResource instance
func NewAPIRefundResource() *APIRefundResource {
	return &APIRefundResource{}
}
