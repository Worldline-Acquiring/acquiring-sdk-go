// This file was automatically generated.

package domain

// APIPaymentRequest represents class ApiPaymentRequest
type APIPaymentRequest struct {
	Amount                    *AmountData        `json:"amount,omitempty"`
	AuthorizationType         *string            `json:"authorizationType,omitempty"`
	CardPaymentData           *CardPaymentData   `json:"cardPaymentData,omitempty"`
	DynamicCurrencyConversion *DccData           `json:"dynamicCurrencyConversion,omitempty"`
	Merchant                  *MerchantData      `json:"merchant,omitempty"`
	OperationID               *string            `json:"operationId,omitempty"`
	References                *PaymentReferences `json:"references,omitempty"`
	TransactionTimestamp      *string            `json:"transactionTimestamp,omitempty"`
}

// NewAPIPaymentRequest constructs a new APIPaymentRequest instance
func NewAPIPaymentRequest() *APIPaymentRequest {
	return &APIPaymentRequest{}
}
