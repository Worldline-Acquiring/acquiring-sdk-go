// This file was automatically generated.

package domain

// PaymentReferences represents class PaymentReferences
type PaymentReferences struct {
	DynamicDescriptor        *string `json:"dynamicDescriptor,omitempty"`
	MerchantReference        *string `json:"merchantReference,omitempty"`
	RetrievalReferenceNumber *string `json:"retrievalReferenceNumber,omitempty"`
}

// NewPaymentReferences constructs a new PaymentReferences instance
func NewPaymentReferences() *PaymentReferences {
	return &PaymentReferences{}
}
