// This file was automatically generated.

package domain

// MerchantData represents class MerchantData
type MerchantData struct {
	Address              *string `json:"address,omitempty"`
	City                 *string `json:"city,omitempty"`
	CountryCode          *string `json:"countryCode,omitempty"`
	MerchantCategoryCode *int32  `json:"merchantCategoryCode,omitempty"`
	Name                 *string `json:"name,omitempty"`
	PostalCode           *string `json:"postalCode,omitempty"`
	StateCode            *string `json:"stateCode,omitempty"`
}

// NewMerchantData constructs a new MerchantData instance
func NewMerchantData() *MerchantData {
	return &MerchantData{}
}
