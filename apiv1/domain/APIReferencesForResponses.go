// This file was automatically generated.

package domain

// APIReferencesForResponses represents class ApiReferencesForResponses.
type APIReferencesForResponses struct {
	PaymentAccountReference  *string `json:"paymentAccountReference,omitempty"`
	RetrievalReferenceNumber *string `json:"retrievalReferenceNumber,omitempty"`
	SchemeTransactionID      *string `json:"schemeTransactionId,omitempty"`
	SchemeTransactionLinkID  *string `json:"schemeTransactionLinkId,omitempty"`
}

// NewAPIReferencesForResponses constructs a new APIReferencesForResponses instance.
func NewAPIReferencesForResponses() *APIReferencesForResponses {
	return &APIReferencesForResponses{}
}
