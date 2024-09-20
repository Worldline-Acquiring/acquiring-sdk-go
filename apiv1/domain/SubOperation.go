// This file was automatically generated.

package domain

// SubOperation represents class SubOperation
type SubOperation struct {
	Amount                  *AmountData `json:"amount,omitempty"`
	AuthorizationCode       *string     `json:"authorizationCode,omitempty"`
	OperationID             *string     `json:"operationId,omitempty"`
	OperationTimestamp      *string     `json:"operationTimestamp,omitempty"`
	OperationType           *string     `json:"operationType,omitempty"`
	ResponseCode            *string     `json:"responseCode,omitempty"`
	ResponseCodeCategory    *string     `json:"responseCodeCategory,omitempty"`
	ResponseCodeDescription *string     `json:"responseCodeDescription,omitempty"`
	RetryAfter              *string     `json:"retryAfter,omitempty"`
}

// NewSubOperation constructs a new SubOperation instance
func NewSubOperation() *SubOperation {
	return &SubOperation{}
}
