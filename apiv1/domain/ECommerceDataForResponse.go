// This file was automatically generated.

package domain

// ECommerceDataForResponse represents class ECommerceDataForResponse
type ECommerceDataForResponse struct {
	AddressVerificationResult *string `json:"addressVerificationResult,omitempty"`
	CardSecurityCodeResult    *string `json:"cardSecurityCodeResult,omitempty"`
}

// NewECommerceDataForResponse constructs a new ECommerceDataForResponse instance
func NewECommerceDataForResponse() *ECommerceDataForResponse {
	return &ECommerceDataForResponse{}
}
