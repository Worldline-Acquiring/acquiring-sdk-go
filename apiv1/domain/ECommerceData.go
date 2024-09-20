// This file was automatically generated.

package domain

// ECommerceData represents class ECommerceData
type ECommerceData struct {
	AddressVerificationData *AddressVerificationData `json:"addressVerificationData,omitempty"`
	ScaExemptionRequest     *string                  `json:"scaExemptionRequest,omitempty"`
	ThreeDSecure            *ThreeDSecure            `json:"threeDSecure,omitempty"`
}

// NewECommerceData constructs a new ECommerceData instance
func NewECommerceData() *ECommerceData {
	return &ECommerceData{}
}
