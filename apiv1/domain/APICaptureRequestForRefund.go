// This file was automatically generated.

package domain

// APICaptureRequestForRefund represents class ApiCaptureRequestForRefund.
type APICaptureRequestForRefund struct {
	OperationID          *string            `json:"operationId,omitempty"`
	References           *PaymentReferences `json:"references,omitempty"`
	TerminalData         *TerminalData      `json:"terminalData,omitempty"`
	TransactionTimestamp *string            `json:"transactionTimestamp,omitempty"`
}

// NewAPICaptureRequestForRefund constructs a new APICaptureRequestForRefund instance.
func NewAPICaptureRequestForRefund() *APICaptureRequestForRefund {
	return &APICaptureRequestForRefund{}
}
