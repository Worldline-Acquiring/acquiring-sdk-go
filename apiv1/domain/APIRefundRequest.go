// This file was automatically generated.

package domain

// APIRefundRequest represents class ApiRefundRequest
type APIRefundRequest struct {
	Amount                    *AmountData               `json:"amount,omitempty"`
	CardPaymentData           *CardPaymentDataForRefund `json:"cardPaymentData,omitempty"`
	DynamicCurrencyConversion *DccData                  `json:"dynamicCurrencyConversion,omitempty"`
	Merchant                  *MerchantData             `json:"merchant,omitempty"`
	OperationID               *string                   `json:"operationId,omitempty"`
	References                *PaymentReferences        `json:"references,omitempty"`
	TransactionTimestamp      *string                   `json:"transactionTimestamp,omitempty"`
}

// NewAPIRefundRequest constructs a new APIRefundRequest instance
func NewAPIRefundRequest() *APIRefundRequest {
	return &APIRefundRequest{}
}
