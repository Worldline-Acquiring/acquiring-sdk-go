// This file was automatically generated.

package domain

// CapturePointOfSaleData represents class CapturePointOfSaleData.
type CapturePointOfSaleData struct {
	EmvData *[]EmvDataItem `json:"emvData,omitempty"`
}

// NewCapturePointOfSaleData constructs a new CapturePointOfSaleData instance.
func NewCapturePointOfSaleData() *CapturePointOfSaleData {
	return &CapturePointOfSaleData{}
}
