// This file was automatically generated.

package domain

// OriginalTransactionData represents class OriginalTransactionData.
type OriginalTransactionData struct {
	CardholderVerificationMethod *string                       `json:"cardholderVerificationMethod,omitempty"`
	LocalTransactionDateTime     *string                       `json:"localTransactionDateTime,omitempty"`
	PaymentCategory              *string                       `json:"paymentCategory,omitempty"`
	PaymentMethodData            *PaymentMethodData            `json:"paymentMethodData,omitempty"`
	PointOfSaleEntryMode         *string                       `json:"pointOfSaleEntryMode,omitempty"`
	SchemeProcessedDateTime      *string                       `json:"schemeProcessedDateTime,omitempty"`
	SettlementAmount             *AmountData                   `json:"settlementAmount,omitempty"`
	TransactionAmount            *AmountData                   `json:"transactionAmount,omitempty"`
	TransactionReferences        *TransactionReferencesDispute `json:"transactionReferences,omitempty"`
}

// NewOriginalTransactionData constructs a new OriginalTransactionData instance.
func NewOriginalTransactionData() *OriginalTransactionData {
	return &OriginalTransactionData{}
}
