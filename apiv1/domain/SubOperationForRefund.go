// This file was automatically generated.

package domain

// SubOperationForRefund represents class SubOperationForRefund
type SubOperationForRefund struct {
	Amount                  *AmountData `json:"amount,omitempty"`
	OperationID             *string     `json:"operationId,omitempty"`
	OperationTimestamp      *string     `json:"operationTimestamp,omitempty"`
	OperationType           *string     `json:"operationType,omitempty"`
	ResponseCode            *string     `json:"responseCode,omitempty"`
	ResponseCodeCategory    *string     `json:"responseCodeCategory,omitempty"`
	ResponseCodeDescription *string     `json:"responseCodeDescription,omitempty"`
	RetryAfter              *string     `json:"retryAfter,omitempty"`
}

// NewSubOperationForRefund constructs a new SubOperationForRefund instance
func NewSubOperationForRefund() *SubOperationForRefund {
	return &SubOperationForRefund{}
}
