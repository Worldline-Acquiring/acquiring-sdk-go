// This file was automatically generated.

package domain

// OnlinePinData represents class OnlinePinData.
type OnlinePinData struct {
	EncryptedPinBlock *string                    `json:"encryptedPinBlock,omitempty"`
	PinBlockFormat    *int32                     `json:"pinBlockFormat,omitempty"`
	PinEncryptionData PinEncryptionDataInterface `json:"pinEncryptionData,omitempty"`
}

// NewOnlinePinData constructs a new OnlinePinData instance.
func NewOnlinePinData() *OnlinePinData {
	return &OnlinePinData{}
}
