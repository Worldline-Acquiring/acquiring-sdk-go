// This file was automatically generated.

package domain

// APITechnicalReversalRequest represents class ApiTechnicalReversalRequest
type APITechnicalReversalRequest struct {
	OperationID          *string `json:"operationId,omitempty"`
	Reason               *string `json:"reason,omitempty"`
	TransactionTimestamp *string `json:"transactionTimestamp,omitempty"`
}

// NewAPITechnicalReversalRequest constructs a new APITechnicalReversalRequest instance
func NewAPITechnicalReversalRequest() *APITechnicalReversalRequest {
	return &APITechnicalReversalRequest{}
}
