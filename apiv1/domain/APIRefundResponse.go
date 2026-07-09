// This file was automatically generated.

package domain

// APIRefundResponse represents class ApiRefundResponse.
type APIRefundResponse struct {
	AdditionalResponseData  *AdditionalResponseData     `json:"additionalResponseData,omitempty"`
	AuthorizationCode       *string                     `json:"authorizationCode,omitempty"`
	CardPaymentData         *CardPaymentDataForResponse `json:"cardPaymentData,omitempty"`
	EmvData                 *[]EmvDataItem              `json:"emvData,omitempty"`
	OperationID             *string                     `json:"operationId,omitempty"`
	ReferencedPaymentID     *string                     `json:"referencedPaymentId,omitempty"`
	References              *APIReferencesForResponses  `json:"references,omitempty"`
	RefundID                *string                     `json:"refundId,omitempty"`
	Responder               *string                     `json:"responder,omitempty"`
	ResponseCode            *string                     `json:"responseCode,omitempty"`
	ResponseCodeCategory    *string                     `json:"responseCodeCategory,omitempty"`
	ResponseCodeDescription *string                     `json:"responseCodeDescription,omitempty"`
	Status                  *string                     `json:"status,omitempty"`
	StatusTimestamp         *string                     `json:"statusTimestamp,omitempty"`
	TotalAuthorizedAmount   *AmountData                 `json:"totalAuthorizedAmount,omitempty"`
}

// NewAPIRefundResponse constructs a new APIRefundResponse instance.
func NewAPIRefundResponse() *APIRefundResponse {
	return &APIRefundResponse{}
}
