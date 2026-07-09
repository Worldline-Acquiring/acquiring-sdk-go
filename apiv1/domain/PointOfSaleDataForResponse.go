// This file was automatically generated.

package domain

// PointOfSaleDataForResponse represents class PointOfSaleDataForResponse.
type PointOfSaleDataForResponse struct {
	EmvData         *[]EmvDataItem `json:"emvData,omitempty"`
	PanLast4Digits  *string        `json:"panLast4Digits,omitempty"`
	PinRetryCounter *int32         `json:"pinRetryCounter,omitempty"`
}

// NewPointOfSaleDataForResponse constructs a new PointOfSaleDataForResponse instance.
func NewPointOfSaleDataForResponse() *PointOfSaleDataForResponse {
	return &PointOfSaleDataForResponse{}
}
