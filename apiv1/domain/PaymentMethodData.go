// This file was automatically generated.

package domain

// PaymentMethodData represents class PaymentMethodData.
type PaymentMethodData struct {
	PaymentMethodDataBase
	IssuingCountryCode *string `json:"issuingCountryCode,omitempty"`
}

// NewPaymentMethodData constructs a new PaymentMethodData instance.
func NewPaymentMethodData() *PaymentMethodData {
	return &PaymentMethodData{}
}
