// This file was automatically generated.

package domain

// APIAccountVerificationResponse represents class ApiAccountVerificationResponse
type APIAccountVerificationResponse struct {
	AuthorizationCode       *string                     `json:"authorizationCode,omitempty"`
	CardPaymentData         *CardPaymentDataForResponse `json:"cardPaymentData,omitempty"`
	OperationID             *string                     `json:"operationId,omitempty"`
	References              *APIReferencesForResponses  `json:"references,omitempty"`
	Responder               *string                     `json:"responder,omitempty"`
	ResponseCode            *string                     `json:"responseCode,omitempty"`
	ResponseCodeCategory    *string                     `json:"responseCodeCategory,omitempty"`
	ResponseCodeDescription *string                     `json:"responseCodeDescription,omitempty"`
}

// NewAPIAccountVerificationResponse constructs a new APIAccountVerificationResponse instance
func NewAPIAccountVerificationResponse() *APIAccountVerificationResponse {
	return &APIAccountVerificationResponse{}
}
