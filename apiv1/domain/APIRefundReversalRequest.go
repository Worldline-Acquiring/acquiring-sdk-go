// This file was automatically generated.

package domain

// APIRefundReversalRequest represents class ApiRefundReversalRequest.
type APIRefundReversalRequest struct {
	OperationID          *string       `json:"operationId,omitempty"`
	TerminalData         *TerminalData `json:"terminalData,omitempty"`
	TransactionTimestamp *string       `json:"transactionTimestamp,omitempty"`
}

// NewAPIRefundReversalRequest constructs a new APIRefundReversalRequest instance.
func NewAPIRefundReversalRequest() *APIRefundReversalRequest {
	return &APIRefundReversalRequest{}
}
