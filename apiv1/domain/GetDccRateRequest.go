// This file was automatically generated.

package domain

// GetDccRateRequest represents class GetDccRateRequest.
type GetDccRateRequest struct {
	CardPaymentData *CardDataForDcc        `json:"cardPaymentData,omitempty"`
	OperationID     *string                `json:"operationId,omitempty"`
	PointOfSaleData *PointOfSaleDataForDcc `json:"pointOfSaleData,omitempty"`
	RateReferenceID *string                `json:"rateReferenceId,omitempty"`
	TargetCurrency  *string                `json:"targetCurrency,omitempty"`
	Transaction     *TransactionDataForDcc `json:"transaction,omitempty"`
}

// NewGetDccRateRequest constructs a new GetDccRateRequest instance.
func NewGetDccRateRequest() *GetDccRateRequest {
	return &GetDccRateRequest{}
}
