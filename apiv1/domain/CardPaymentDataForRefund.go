// This file was automatically generated.

package domain

// CardPaymentDataForRefund represents class CardPaymentDataForRefund
type CardPaymentDataForRefund struct {
	Brand              *string           `json:"brand,omitempty"`
	CaptureImmediately *bool             `json:"captureImmediately,omitempty"`
	CardData           *PlainCardData    `json:"cardData,omitempty"`
	CardEntryMode      *string           `json:"cardEntryMode,omitempty"`
	NetworkTokenData   *NetworkTokenData `json:"networkTokenData,omitempty"`
	PointOfSaleData    *PointOfSaleData  `json:"pointOfSaleData,omitempty"`
	WalletID           *string           `json:"walletId,omitempty"`
}

// NewCardPaymentDataForRefund constructs a new CardPaymentDataForRefund instance
func NewCardPaymentDataForRefund() *CardPaymentDataForRefund {
	return &CardPaymentDataForRefund{}
}
