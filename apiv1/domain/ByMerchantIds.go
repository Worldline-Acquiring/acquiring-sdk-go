// This file was automatically generated.

package domain

import "encoding/json"

// ByMerchantIdsMerchantScopeType is the fixed merchantScopeType of ByMerchantIds.
const ByMerchantIdsMerchantScopeType = "BY_MERCHANT_IDS"

// ByMerchantIds represents class ByMerchantIds.
type ByMerchantIds struct {
	MerchantScope
	MerchantIds *[]MerchantIDItem `json:"merchantIds,omitempty"`
}

// NewByMerchantIds constructs a new ByMerchantIds instance.
func NewByMerchantIds() *ByMerchantIds {
	value := ByMerchantIdsMerchantScopeType
	return &ByMerchantIds{
		MerchantScope: MerchantScope{
			MerchantScopeType: &value,
		},
	}
}

// MarshalJSON implements json.Marshaler. It pins merchantScopeType to
// ByMerchantIdsMerchantScopeType, whatever the embedded MerchantScope holds.
func (o ByMerchantIds) MarshalJSON() ([]byte, error) {
	type alias ByMerchantIds

	return json.Marshal(struct {
		MerchantScopeType string `json:"merchantScopeType"`
		alias
	}{
		MerchantScopeType: ByMerchantIdsMerchantScopeType,
		alias:             alias(o),
	})
}
