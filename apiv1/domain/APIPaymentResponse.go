// This file was automatically generated.

package domain

// APIPaymentResponse represents class ApiPaymentResponse
type APIPaymentResponse struct {
	CardPaymentData          *CardPaymentDataForResponse `json:"cardPaymentData,omitempty"`
	InitialAuthorizationCode *string                     `json:"initialAuthorizationCode,omitempty"`
	OperationID              *string                     `json:"operationId,omitempty"`
	PaymentID                *string                     `json:"paymentId,omitempty"`
	References               *APIReferencesForResponses  `json:"references,omitempty"`
	Responder                *string                     `json:"responder,omitempty"`
	ResponseCode             *string                     `json:"responseCode,omitempty"`
	ResponseCodeCategory     *string                     `json:"responseCodeCategory,omitempty"`
	ResponseCodeDescription  *string                     `json:"responseCodeDescription,omitempty"`
	RetryAfter               *string                     `json:"retryAfter,omitempty"`
	Status                   *string                     `json:"status,omitempty"`
	StatusTimestamp          *string                     `json:"statusTimestamp,omitempty"`
	TotalAuthorizedAmount    *AmountData                 `json:"totalAuthorizedAmount,omitempty"`
}

// NewAPIPaymentResponse constructs a new APIPaymentResponse instance
func NewAPIPaymentResponse() *APIPaymentResponse {
	return &APIPaymentResponse{}
}
