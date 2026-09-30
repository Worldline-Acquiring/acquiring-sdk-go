// This file was automatically generated.

package domain

import "encoding/json"

// ByMerchantRootIdsMerchantScopeType is the fixed merchantScopeType of ByMerchantRootIds.
const ByMerchantRootIdsMerchantScopeType = "BY_MERCHANT_ROOT_IDS"

// ByMerchantRootIds represents class ByMerchantRootIds.
type ByMerchantRootIds struct {
	MerchantScope
	MerchantRootIds *[]MerchantRootIDItem `json:"merchantRootIds,omitempty"`
}

// NewByMerchantRootIds constructs a new ByMerchantRootIds instance.
func NewByMerchantRootIds() *ByMerchantRootIds {
	value := ByMerchantRootIdsMerchantScopeType
	return &ByMerchantRootIds{
		MerchantScope: MerchantScope{
			MerchantScopeType: &value,
		},
	}
}

// MarshalJSON implements json.Marshaler. It pins merchantScopeType to
// ByMerchantRootIdsMerchantScopeType, whatever the embedded MerchantScope holds.
func (o ByMerchantRootIds) MarshalJSON() ([]byte, error) {
	type alias ByMerchantRootIds

	return json.Marshal(struct {
		MerchantScopeType string `json:"merchantScopeType"`
		alias
	}{
		MerchantScopeType: ByMerchantRootIdsMerchantScopeType,
		alias:             alias(o),
	})
}
