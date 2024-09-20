// This file was automatically generated.

package domain

// PlainCardData represents class PlainCardData
type PlainCardData struct {
	CardNumber       *string `json:"cardNumber,omitempty"`
	CardSecurityCode *string `json:"cardSecurityCode,omitempty"`
	ExpiryDate       *string `json:"expiryDate,omitempty"`
}

// NewPlainCardData constructs a new PlainCardData instance
func NewPlainCardData() *PlainCardData {
	return &PlainCardData{}
}
