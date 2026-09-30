// This file was automatically generated.

package domain

// DisputeMerchantData represents class DisputeMerchantData.
type DisputeMerchantData struct {
	DisputeMerchantDataBase
	MerchantCategoryCode *int32  `json:"merchantCategoryCode,omitempty"`
	MerchantCity         *string `json:"merchantCity,omitempty"`
	MerchantCountryCode  *string `json:"merchantCountryCode,omitempty"`
	MerchantName         *string `json:"merchantName,omitempty"`
}

// NewDisputeMerchantData constructs a new DisputeMerchantData instance.
func NewDisputeMerchantData() *DisputeMerchantData {
	return &DisputeMerchantData{}
}
