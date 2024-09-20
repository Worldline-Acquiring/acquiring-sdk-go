// This file was automatically generated.

package domain

// DccData represents class DccData
type DccData struct {
	Amount           *int64   `json:"amount,omitempty"`
	ConversionRate   *float64 `json:"conversionRate,omitempty"`
	CurrencyCode     *string  `json:"currencyCode,omitempty"`
	NumberOfDecimals *int32   `json:"numberOfDecimals,omitempty"`
}

// NewDccData constructs a new DccData instance
func NewDccData() *DccData {
	return &DccData{}
}
