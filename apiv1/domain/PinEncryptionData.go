// This file was automatically generated.

package domain

// PinEncryptionData represents class PinEncryptionData.
type PinEncryptionData struct {
	PinEncryptionType *string `json:"pinEncryptionType,omitempty"`
}

// NewPinEncryptionData constructs a new PinEncryptionData instance.
func NewPinEncryptionData() *PinEncryptionData {
	return &PinEncryptionData{}
}

// isPinEncryptionData marks PinEncryptionData and every schema that extends it with allOf.
func (o PinEncryptionData) isPinEncryptionData() {}

// PinEncryptionDataInterface is implemented by PinEncryptionData and by every schema
// that extends it with allOf.
type PinEncryptionDataInterface interface {
	isPinEncryptionData()
}
