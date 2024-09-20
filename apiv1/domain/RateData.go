// This file was automatically generated.

package domain

// RateData represents class RateData
type RateData struct {
	ExchangeRate         *float64 `json:"exchangeRate,omitempty"`
	InvertedExchangeRate *float64 `json:"invertedExchangeRate,omitempty"`
	MarkUp               *float64 `json:"markUp,omitempty"`
	MarkUpBasis          *string  `json:"markUpBasis,omitempty"`
	QuotationDateTime    *string  `json:"quotationDateTime,omitempty"`
}

// NewRateData constructs a new RateData instance
func NewRateData() *RateData {
	return &RateData{}
}
