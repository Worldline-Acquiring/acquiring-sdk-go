// This file was automatically generated.

package domain

// PaymentMethodDataBase represents class PaymentMethodDataBase.
type PaymentMethodDataBase struct {
	MaskedIdentifier *string `json:"maskedIdentifier,omitempty"`
	Scheme           *string `json:"scheme,omitempty"`
	SchemeBrand      *string `json:"schemeBrand,omitempty"`
}

// NewPaymentMethodDataBase constructs a new PaymentMethodDataBase instance.
func NewPaymentMethodDataBase() *PaymentMethodDataBase {
	return &PaymentMethodDataBase{}
}
