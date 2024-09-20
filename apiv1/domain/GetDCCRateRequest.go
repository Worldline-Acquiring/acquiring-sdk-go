// This file was automatically generated.

package domain

// GetDCCRateRequest represents class GetDCCRateRequest
type GetDCCRateRequest struct {
	CardPaymentData *CardDataForDcc        `json:"cardPaymentData,omitempty"`
	OperationID     *string                `json:"operationId,omitempty"`
	PointOfSaleData *PointOfSaleDataForDcc `json:"pointOfSaleData,omitempty"`
	RateReferenceID *string                `json:"rateReferenceId,omitempty"`
	TargetCurrency  *string                `json:"targetCurrency,omitempty"`
	Transaction     *TransactionDataForDcc `json:"transaction,omitempty"`
}

// NewGetDCCRateRequest constructs a new GetDCCRateRequest instance
func NewGetDCCRateRequest() *GetDCCRateRequest {
	return &GetDCCRateRequest{}
}
