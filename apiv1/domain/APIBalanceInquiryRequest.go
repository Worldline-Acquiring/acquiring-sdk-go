// This file was automatically generated.

package domain

// APIBalanceInquiryRequest represents class ApiBalanceInquiryRequest.
type APIBalanceInquiryRequest struct {
	CardPaymentData      *CardPaymentDataForBalanceInquiry `json:"cardPaymentData,omitempty"`
	Merchant             *MerchantData                     `json:"merchant,omitempty"`
	OperationID          *string                           `json:"operationId,omitempty"`
	References           *PaymentReferences                `json:"references,omitempty"`
	TerminalData         *TerminalData                     `json:"terminalData,omitempty"`
	TransactionTimestamp *string                           `json:"transactionTimestamp,omitempty"`
}

// NewAPIBalanceInquiryRequest constructs a new APIBalanceInquiryRequest instance.
func NewAPIBalanceInquiryRequest() *APIBalanceInquiryRequest {
	return &APIBalanceInquiryRequest{}
}
