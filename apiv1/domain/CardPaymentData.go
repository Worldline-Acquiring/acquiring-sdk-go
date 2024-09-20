// This file was automatically generated.

package domain

// CardPaymentData represents class CardPaymentData
type CardPaymentData struct {
	AllowPartialApproval         *bool             `json:"allowPartialApproval,omitempty"`
	Brand                        *string           `json:"brand,omitempty"`
	CaptureImmediately           *bool             `json:"captureImmediately,omitempty"`
	CardData                     *PlainCardData    `json:"cardData,omitempty"`
	CardEntryMode                *string           `json:"cardEntryMode,omitempty"`
	CardOnFileData               *CardOnFileData   `json:"cardOnFileData,omitempty"`
	CardholderVerificationMethod *string           `json:"cardholderVerificationMethod,omitempty"`
	EcommerceData                *ECommerceData    `json:"ecommerceData,omitempty"`
	NetworkTokenData             *NetworkTokenData `json:"networkTokenData,omitempty"`
	PointOfSaleData              *PointOfSaleData  `json:"pointOfSaleData,omitempty"`
	WalletID                     *string           `json:"walletId,omitempty"`
}

// NewCardPaymentData constructs a new CardPaymentData instance
func NewCardPaymentData() *CardPaymentData {
	return &CardPaymentData{}
}
