// This file was automatically generated.

package domain

// AddressVerificationData represents class AddressVerificationData
type AddressVerificationData struct {
	CardholderAddress    *string `json:"cardholderAddress,omitempty"`
	CardholderPostalCode *string `json:"cardholderPostalCode,omitempty"`
}

// NewAddressVerificationData constructs a new AddressVerificationData instance
func NewAddressVerificationData() *AddressVerificationData {
	return &AddressVerificationData{}
}
