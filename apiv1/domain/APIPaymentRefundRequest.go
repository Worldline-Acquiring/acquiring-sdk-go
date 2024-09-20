// This file was automatically generated.

package domain

// APIPaymentRefundRequest represents class ApiPaymentRefundRequest
type APIPaymentRefundRequest struct {
	Amount                    *AmountData        `json:"amount,omitempty"`
	CaptureImmediately        *bool              `json:"captureImmediately,omitempty"`
	DynamicCurrencyConversion *DccData           `json:"dynamicCurrencyConversion,omitempty"`
	OperationID               *string            `json:"operationId,omitempty"`
	References                *PaymentReferences `json:"references,omitempty"`
	TransactionTimestamp      *string            `json:"transactionTimestamp,omitempty"`
}

// NewAPIPaymentRefundRequest constructs a new APIPaymentRefundRequest instance
func NewAPIPaymentRefundRequest() *APIPaymentRefundRequest {
	return &APIPaymentRefundRequest{}
}
