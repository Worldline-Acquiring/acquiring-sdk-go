// This file was automatically generated.

package domain

// DisputeMerchantDataBase represents class DisputeMerchantDataBase.
type DisputeMerchantDataBase struct {
	AcquirerID     *string `json:"acquirerId,omitempty"`
	MerchantID     *string `json:"merchantId,omitempty"`
	MerchantRootID *string `json:"merchantRootId,omitempty"`
}

// NewDisputeMerchantDataBase constructs a new DisputeMerchantDataBase instance.
func NewDisputeMerchantDataBase() *DisputeMerchantDataBase {
	return &DisputeMerchantDataBase{}
}
