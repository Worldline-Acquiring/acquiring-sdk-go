// This file was automatically generated.

package domain

// ECommerceDataForAccountVerification represents class ECommerceDataForAccountVerification
type ECommerceDataForAccountVerification struct {
	AddressVerificationData *AddressVerificationData `json:"addressVerificationData,omitempty"`
	ThreeDSecure            *ThreeDSecure            `json:"threeDSecure,omitempty"`
}

// NewECommerceDataForAccountVerification constructs a new ECommerceDataForAccountVerification instance
func NewECommerceDataForAccountVerification() *ECommerceDataForAccountVerification {
	return &ECommerceDataForAccountVerification{}
}
