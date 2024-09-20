// This file was automatically generated.

package domain

// APIPaymentReversalRequest represents class ApiPaymentReversalRequest
type APIPaymentReversalRequest struct {
	DynamicCurrencyConversion *DccData    `json:"dynamicCurrencyConversion,omitempty"`
	OperationID               *string     `json:"operationId,omitempty"`
	ReversalAmount            *AmountData `json:"reversalAmount,omitempty"`
	TransactionTimestamp      *string     `json:"transactionTimestamp,omitempty"`
}

// NewAPIPaymentReversalRequest constructs a new APIPaymentReversalRequest instance
func NewAPIPaymentReversalRequest() *APIPaymentReversalRequest {
	return &APIPaymentReversalRequest{}
}
