// This file was automatically generated.

package domain

// CardPaymentDataForVerification represents class CardPaymentDataForVerification
type CardPaymentDataForVerification struct {
	Brand                        *string                              `json:"brand,omitempty"`
	CardData                     *PlainCardData                       `json:"cardData,omitempty"`
	CardEntryMode                *string                              `json:"cardEntryMode,omitempty"`
	CardOnFileData               *CardOnFileData                      `json:"cardOnFileData,omitempty"`
	CardholderVerificationMethod *string                              `json:"cardholderVerificationMethod,omitempty"`
	EcommerceData                *ECommerceDataForAccountVerification `json:"ecommerceData,omitempty"`
	NetworkTokenData             *NetworkTokenData                    `json:"networkTokenData,omitempty"`
	WalletID                     *string                              `json:"walletId,omitempty"`
}

// NewCardPaymentDataForVerification constructs a new CardPaymentDataForVerification instance
func NewCardPaymentDataForVerification() *CardPaymentDataForVerification {
	return &CardPaymentDataForVerification{}
}
