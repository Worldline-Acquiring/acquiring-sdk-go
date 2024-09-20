// This file was automatically generated.

package domain

// APIIncrementRequest represents class ApiIncrementRequest
type APIIncrementRequest struct {
	DynamicCurrencyConversion *DccData    `json:"dynamicCurrencyConversion,omitempty"`
	IncrementAmount           *AmountData `json:"incrementAmount,omitempty"`
	OperationID               *string     `json:"operationId,omitempty"`
	TransactionTimestamp      *string     `json:"transactionTimestamp,omitempty"`
}

// NewAPIIncrementRequest constructs a new APIIncrementRequest instance
func NewAPIIncrementRequest() *APIIncrementRequest {
	return &APIIncrementRequest{}
}
