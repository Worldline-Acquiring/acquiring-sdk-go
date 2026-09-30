// This file was automatically generated.

package domain

// TransactionReferencesDispute represents class TransactionReferencesDispute.
type TransactionReferencesDispute struct {
	TransactionReferencesBase
	AcquirerTransactionReference *string `json:"acquirerTransactionReference,omitempty"`
	MerchantOperationID          *string `json:"merchantOperationId,omitempty"`
	RetrievalReferenceNumber     *string `json:"retrievalReferenceNumber,omitempty"`
	SchemeTransactionID          *string `json:"schemeTransactionId,omitempty"`
	SchemeTransactionLinkID      *string `json:"schemeTransactionLinkId,omitempty"`
	TerminalID                   *string `json:"terminalId,omitempty"`
	TerminalTransactionReference *string `json:"terminalTransactionReference,omitempty"`
}

// NewTransactionReferencesDispute constructs a new TransactionReferencesDispute instance.
func NewTransactionReferencesDispute() *TransactionReferencesDispute {
	return &TransactionReferencesDispute{}
}
