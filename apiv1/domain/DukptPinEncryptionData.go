// This file was automatically generated.

package domain

import "encoding/json"

// DukptPinEncryptionDataPinEncryptionType is the fixed pinEncryptionType of DukptPinEncryptionData.
const DukptPinEncryptionDataPinEncryptionType = "DUKPT"

// DukptPinEncryptionData represents class DukptPinEncryptionData.
type DukptPinEncryptionData struct {
	PinEncryptionData
	KeySerialNumber *string `json:"keySerialNumber,omitempty"`
}

// NewDukptPinEncryptionData constructs a new DukptPinEncryptionData instance.
func NewDukptPinEncryptionData() *DukptPinEncryptionData {
	value := DukptPinEncryptionDataPinEncryptionType
	return &DukptPinEncryptionData{
		PinEncryptionData: PinEncryptionData{
			PinEncryptionType: &value,
		},
	}
}

// MarshalJSON implements json.Marshaler. It pins pinEncryptionType to
// DukptPinEncryptionDataPinEncryptionType, whatever the embedded PinEncryptionData holds.
func (o DukptPinEncryptionData) MarshalJSON() ([]byte, error) {
	type alias DukptPinEncryptionData

	return json.Marshal(struct {
		PinEncryptionType string `json:"pinEncryptionType"`
		alias
	}{
		PinEncryptionType: DukptPinEncryptionDataPinEncryptionType,
		alias:             alias(o),
	})
}
