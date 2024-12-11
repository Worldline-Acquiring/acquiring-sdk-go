// This file was automatically generated.

package domain

// CardPaymentDataForBalanceInquiry represents class CardPaymentDataForBalanceInquiry
type CardPaymentDataForBalanceInquiry struct {
	Brand                        *string          `json:"brand,omitempty"`
	BrandSelector                *string          `json:"brandSelector,omitempty"`
	CardData                     *PlainCardData   `json:"cardData,omitempty"`
	CardEntryMode                *string          `json:"cardEntryMode,omitempty"`
	CardholderVerificationMethod *string          `json:"cardholderVerificationMethod,omitempty"`
	EcommerceData                *ECommerceData   `json:"ecommerceData,omitempty"`
	PointOfSaleData              *PointOfSaleData `json:"pointOfSaleData,omitempty"`
	WalletID                     *string          `json:"walletId,omitempty"`
}

// NewCardPaymentDataForBalanceInquiry constructs a new CardPaymentDataForBalanceInquiry instance
func NewCardPaymentDataForBalanceInquiry() *CardPaymentDataForBalanceInquiry {
	return &CardPaymentDataForBalanceInquiry{}
}
