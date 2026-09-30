// This file was automatically generated.

package domain

import "encoding/json"

// AesUkptPinEncryptionDataPinEncryptionType is the fixed pinEncryptionType of AesUkptPinEncryptionData.
const AesUkptPinEncryptionDataPinEncryptionType = "AES_UKPT"

// AesUkptPinEncryptionData represents class AesUkptPinEncryptionData.
type AesUkptPinEncryptionData struct {
	PinEncryptionData
	KeyGeneration *int32  `json:"keyGeneration,omitempty"`
	RandomValue   *string `json:"randomValue,omitempty"`
}

// NewAesUkptPinEncryptionData constructs a new AesUkptPinEncryptionData instance.
func NewAesUkptPinEncryptionData() *AesUkptPinEncryptionData {
	value := AesUkptPinEncryptionDataPinEncryptionType
	return &AesUkptPinEncryptionData{
		PinEncryptionData: PinEncryptionData{
			PinEncryptionType: &value,
		},
	}
}

// MarshalJSON implements json.Marshaler. It pins pinEncryptionType to
// AesUkptPinEncryptionDataPinEncryptionType, whatever the embedded PinEncryptionData holds.
func (o AesUkptPinEncryptionData) MarshalJSON() ([]byte, error) {
	type alias AesUkptPinEncryptionData

	return json.Marshal(struct {
		PinEncryptionType string `json:"pinEncryptionType"`
		alias
	}{
		PinEncryptionType: AesUkptPinEncryptionDataPinEncryptionType,
		alias:             alias(o),
	})
}
