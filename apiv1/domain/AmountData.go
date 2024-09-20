// This file was automatically generated.

package domain

// AmountData represents class AmountData
type AmountData struct {
	Amount           *int64  `json:"amount,omitempty"`
	CurrencyCode     *string `json:"currencyCode,omitempty"`
	NumberOfDecimals *int32  `json:"numberOfDecimals,omitempty"`
}

// NewAmountData constructs a new AmountData instance
func NewAmountData() *AmountData {
	return &AmountData{}
}
