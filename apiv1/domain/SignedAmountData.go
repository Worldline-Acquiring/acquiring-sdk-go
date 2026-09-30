// This file was automatically generated.

package domain

// SignedAmountData represents class SignedAmountData.
type SignedAmountData struct {
	Amount               *int64  `json:"amount,omitempty"`
	CurrencyCode         *string `json:"currencyCode,omitempty"`
	DebitCreditIndicator *string `json:"debitCreditIndicator,omitempty"`
	NumberOfDecimals     *int32  `json:"numberOfDecimals,omitempty"`
}

// NewSignedAmountData constructs a new SignedAmountData instance.
func NewSignedAmountData() *SignedAmountData {
	return &SignedAmountData{}
}
