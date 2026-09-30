// This file was automatically generated.

package domain

// TransactionReferencesBase represents class TransactionReferencesBase.
type TransactionReferencesBase struct {
	AcquirerReferenceNumber *string `json:"acquirerReferenceNumber,omitempty"`
	MerchantReference       *string `json:"merchantReference,omitempty"`
	PaymentID               *string `json:"paymentId,omitempty"`
}

// NewTransactionReferencesBase constructs a new TransactionReferencesBase instance.
func NewTransactionReferencesBase() *TransactionReferencesBase {
	return &TransactionReferencesBase{}
}
