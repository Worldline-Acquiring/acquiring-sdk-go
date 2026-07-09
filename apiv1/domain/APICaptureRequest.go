// This file was automatically generated.

package domain

// APICaptureRequest represents class ApiCaptureRequest.
type APICaptureRequest struct {
	Amount                     *AmountData                 `json:"amount,omitempty"`
	CaptureAmountBreakdownData *CaptureAmountBreakdownData `json:"captureAmountBreakdownData,omitempty"`
	CaptureSequenceNumber      *int32                      `json:"captureSequenceNumber,omitempty"`
	DynamicCurrencyConversion  *DccData                    `json:"dynamicCurrencyConversion,omitempty"`
	IsFinal                    *bool                       `json:"isFinal,omitempty"`
	MarketplaceData            *MarketplaceData            `json:"marketplaceData,omitempty"`
	OperationID                *string                     `json:"operationId,omitempty"`
	References                 *PaymentReferences          `json:"references,omitempty"`
	TerminalData               *TerminalData               `json:"terminalData,omitempty"`
	TransactionTimestamp       *string                     `json:"transactionTimestamp,omitempty"`
}

// NewAPICaptureRequest constructs a new APICaptureRequest instance.
func NewAPICaptureRequest() *APICaptureRequest {
	return &APICaptureRequest{}
}
