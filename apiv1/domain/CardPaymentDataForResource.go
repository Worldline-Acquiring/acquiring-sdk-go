// This file was automatically generated.

package domain

// CardPaymentDataForResource represents class CardPaymentDataForResource
type CardPaymentDataForResource struct {
	Brand           *string          `json:"brand,omitempty"`
	PointOfSaleData *PointOfSaleData `json:"pointOfSaleData,omitempty"`
}

// NewCardPaymentDataForResource constructs a new CardPaymentDataForResource instance
func NewCardPaymentDataForResource() *CardPaymentDataForResource {
	return &CardPaymentDataForResource{}
}
