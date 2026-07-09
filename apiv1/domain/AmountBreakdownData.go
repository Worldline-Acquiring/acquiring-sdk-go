// This file was automatically generated.

package domain

// AmountBreakdownData represents class AmountBreakdownData.
type AmountBreakdownData struct {
	CashbackAmount *AmountData `json:"cashbackAmount,omitempty"`
	TipAmount      *AmountData `json:"tipAmount,omitempty"`
}

// NewAmountBreakdownData constructs a new AmountBreakdownData instance.
func NewAmountBreakdownData() *AmountBreakdownData {
	return &AmountBreakdownData{}
}
