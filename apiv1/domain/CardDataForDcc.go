// This file was automatically generated.

package domain

// CardDataForDcc represents class CardDataForDcc
type CardDataForDcc struct {
	Bin             *string `json:"bin,omitempty"`
	Brand           *string `json:"brand,omitempty"`
	CardCountryCode *string `json:"cardCountryCode,omitempty"`
	CardEntryMode   *string `json:"cardEntryMode,omitempty"`
}

// NewCardDataForDcc constructs a new CardDataForDcc instance
func NewCardDataForDcc() *CardDataForDcc {
	return &CardDataForDcc{}
}
