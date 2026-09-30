// This file was automatically generated.

package domain

// PointOfSaleData represents class PointOfSaleData.
type PointOfSaleData struct {
	EmvData                       *[]EmvDataItem `json:"emvData,omitempty"`
	IsResponseToPinRequest        *bool          `json:"isResponseToPinRequest,omitempty"`
	IsRetryWithTheSameOperationID *bool          `json:"isRetryWithTheSameOperationId,omitempty"`
	OnlinePinData                 *OnlinePinData `json:"onlinePinData,omitempty"`
	Track2Data                    *string        `json:"track2Data,omitempty"`
}

// NewPointOfSaleData constructs a new PointOfSaleData instance.
func NewPointOfSaleData() *PointOfSaleData {
	return &PointOfSaleData{}
}
