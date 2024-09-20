// This file was automatically generated.

package domain

// TransactionDataForDcc represents class TransactionDataForDcc
type TransactionDataForDcc struct {
	Amount               *AmountData `json:"amount,omitempty"`
	TransactionTimestamp *string     `json:"transactionTimestamp,omitempty"`
	TransactionType      *string     `json:"transactionType,omitempty"`
}

// NewTransactionDataForDcc constructs a new TransactionDataForDcc instance
func NewTransactionDataForDcc() *TransactionDataForDcc {
	return &TransactionDataForDcc{}
}
