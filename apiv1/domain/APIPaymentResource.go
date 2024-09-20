// This file was automatically generated.

package domain

// APIPaymentResource represents class ApiPaymentResource
type APIPaymentResource struct {
	CardPaymentData          *CardPaymentDataForResource `json:"cardPaymentData,omitempty"`
	InitialAuthorizationCode *string                     `json:"initialAuthorizationCode,omitempty"`
	Operations               *[]SubOperation             `json:"operations,omitempty"`
	PaymentID                *string                     `json:"paymentId,omitempty"`
	References               *APIReferencesForResponses  `json:"references,omitempty"`
	RetryAfter               *string                     `json:"retryAfter,omitempty"`
	Status                   *string                     `json:"status,omitempty"`
	StatusTimestamp          *string                     `json:"statusTimestamp,omitempty"`
	TotalAuthorizedAmount    *AmountData                 `json:"totalAuthorizedAmount,omitempty"`
}

// NewAPIPaymentResource constructs a new APIPaymentResource instance
func NewAPIPaymentResource() *APIPaymentResource {
	return &APIPaymentResource{}
}
