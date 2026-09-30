// This file was automatically generated.

package domain

// MerchantScope represents class MerchantScope.
type MerchantScope struct {
	MerchantScopeType *string `json:"merchantScopeType,omitempty"`
}

// NewMerchantScope constructs a new MerchantScope instance.
func NewMerchantScope() *MerchantScope {
	return &MerchantScope{}
}

// isMerchantScope marks MerchantScope and every schema that extends it with allOf.
func (o MerchantScope) isMerchantScope() {}

// MerchantScopeInterface is implemented by MerchantScope and by every schema
// that extends it with allOf.
type MerchantScopeInterface interface {
	isMerchantScope()
}
