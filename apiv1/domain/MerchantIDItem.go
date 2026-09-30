// This file was automatically generated.

package domain

// MerchantIDItem represents class MerchantIdItem.
type MerchantIDItem struct {
	AcquirerID *string `json:"acquirerId,omitempty"`
	MerchantID *string `json:"merchantId,omitempty"`
}

// NewMerchantIDItem constructs a new MerchantIDItem instance.
func NewMerchantIDItem() *MerchantIDItem {
	return &MerchantIDItem{}
}
