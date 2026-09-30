// This file was automatically generated.

package domain

import "encoding/json"

// ZpkPinEncryptionDataPinEncryptionType is the fixed pinEncryptionType of ZpkPinEncryptionData.
const ZpkPinEncryptionDataPinEncryptionType = "ZPK"

// ZpkPinEncryptionData represents class ZpkPinEncryptionData.
type ZpkPinEncryptionData struct {
	PinEncryptionData
	ZonePinKeyID *string `json:"zonePinKeyId,omitempty"`
}

// NewZpkPinEncryptionData constructs a new ZpkPinEncryptionData instance.
func NewZpkPinEncryptionData() *ZpkPinEncryptionData {
	value := ZpkPinEncryptionDataPinEncryptionType
	return &ZpkPinEncryptionData{
		PinEncryptionData: PinEncryptionData{
			PinEncryptionType: &value,
		},
	}
}

// MarshalJSON implements json.Marshaler. It pins pinEncryptionType to
// ZpkPinEncryptionDataPinEncryptionType, whatever the embedded PinEncryptionData holds.
func (o ZpkPinEncryptionData) MarshalJSON() ([]byte, error) {
	type alias ZpkPinEncryptionData

	return json.Marshal(struct {
		PinEncryptionType string `json:"pinEncryptionType"`
		alias
	}{
		PinEncryptionType: ZpkPinEncryptionDataPinEncryptionType,
		alias:             alias(o),
	})
}
