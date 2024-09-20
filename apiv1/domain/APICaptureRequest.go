// This file was automatically generated.

package domain

// APICaptureRequest represents class ApiCaptureRequest
type APICaptureRequest struct {
	Amount                    *AmountData `json:"amount,omitempty"`
	CaptureSequenceNumber     *int32      `json:"captureSequenceNumber,omitempty"`
	DynamicCurrencyConversion *DccData    `json:"dynamicCurrencyConversion,omitempty"`
	IsFinal                   *bool       `json:"isFinal,omitempty"`
	OperationID               *string     `json:"operationId,omitempty"`
	TransactionTimestamp      *string     `json:"transactionTimestamp,omitempty"`
}

// NewAPICaptureRequest constructs a new APICaptureRequest instance
func NewAPICaptureRequest() *APICaptureRequest {
	return &APICaptureRequest{}
}
