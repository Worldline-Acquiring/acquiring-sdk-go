// This file was automatically generated.

package domain

// APIAccountVerificationRequest represents class ApiAccountVerificationRequest
type APIAccountVerificationRequest struct {
	CardPaymentData      *CardPaymentDataForVerification `json:"cardPaymentData,omitempty"`
	Merchant             *MerchantData                   `json:"merchant,omitempty"`
	OperationID          *string                         `json:"operationId,omitempty"`
	References           *PaymentReferences              `json:"references,omitempty"`
	TransactionTimestamp *string                         `json:"transactionTimestamp,omitempty"`
}

// NewAPIAccountVerificationRequest constructs a new APIAccountVerificationRequest instance
func NewAPIAccountVerificationRequest() *APIAccountVerificationRequest {
	return &APIAccountVerificationRequest{}
}
