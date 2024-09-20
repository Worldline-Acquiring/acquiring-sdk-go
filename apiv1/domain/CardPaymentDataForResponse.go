// This file was automatically generated.

package domain

// CardPaymentDataForResponse represents class CardPaymentDataForResponse
type CardPaymentDataForResponse struct {
	Brand           *string                   `json:"brand,omitempty"`
	EcommerceData   *ECommerceDataForResponse `json:"ecommerceData,omitempty"`
	PointOfSaleData *PointOfSaleData          `json:"pointOfSaleData,omitempty"`
}

// NewCardPaymentDataForResponse constructs a new CardPaymentDataForResponse instance
func NewCardPaymentDataForResponse() *CardPaymentDataForResponse {
	return &CardPaymentDataForResponse{}
}
