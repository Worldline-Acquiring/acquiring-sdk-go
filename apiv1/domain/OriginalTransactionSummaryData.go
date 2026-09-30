// This file was automatically generated.

package domain

// OriginalTransactionSummaryData represents class OriginalTransactionSummaryData.
type OriginalTransactionSummaryData struct {
	CardholderVerificationMethod *string                    `json:"cardholderVerificationMethod,omitempty"`
	PaymentCategory              *string                    `json:"paymentCategory,omitempty"`
	PaymentMethodData            *PaymentMethodDataBase     `json:"paymentMethodData,omitempty"`
	PointOfSaleEntryMode         *string                    `json:"pointOfSaleEntryMode,omitempty"`
	TransactionReferences        *TransactionReferencesBase `json:"transactionReferences,omitempty"`
}

// NewOriginalTransactionSummaryData constructs a new OriginalTransactionSummaryData instance.
func NewOriginalTransactionSummaryData() *OriginalTransactionSummaryData {
	return &OriginalTransactionSummaryData{}
}
