// This file was automatically generated.

package domain

import "encoding/json"

// ByAcquirerIdsMerchantScopeType is the fixed merchantScopeType of ByAcquirerIds.
const ByAcquirerIdsMerchantScopeType = "BY_ACQUIRER_IDS"

// ByAcquirerIds represents class ByAcquirerIds.
type ByAcquirerIds struct {
	MerchantScope
	AcquirerIds *[]string `json:"acquirerIds,omitempty"`
}

// NewByAcquirerIds constructs a new ByAcquirerIds instance.
func NewByAcquirerIds() *ByAcquirerIds {
	value := ByAcquirerIdsMerchantScopeType
	return &ByAcquirerIds{
		MerchantScope: MerchantScope{
			MerchantScopeType: &value,
		},
	}
}

// MarshalJSON implements json.Marshaler. It pins merchantScopeType to
// ByAcquirerIdsMerchantScopeType, whatever the embedded MerchantScope holds.
func (o ByAcquirerIds) MarshalJSON() ([]byte, error) {
	type alias ByAcquirerIds

	return json.Marshal(struct {
		MerchantScopeType string `json:"merchantScopeType"`
		alias
	}{
		MerchantScopeType: ByAcquirerIdsMerchantScopeType,
		alias:             alias(o),
	})
}
