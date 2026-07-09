// This file was automatically generated.

package domain

// APIIncrementResponse represents class ApiIncrementResponse.
type APIIncrementResponse struct {
	AdditionalResponseData  *AdditionalResponseData       `json:"additionalResponseData,omitempty"`
	AuthorizationCode       *string                       `json:"authorizationCode,omitempty"`
	OperationID             *string                       `json:"operationId,omitempty"`
	Payment                 *APIPaymentSummaryForResponse `json:"payment,omitempty"`
	Responder               *string                       `json:"responder,omitempty"`
	ResponseCode            *string                       `json:"responseCode,omitempty"`
	ResponseCodeCategory    *string                       `json:"responseCodeCategory,omitempty"`
	ResponseCodeDescription *string                       `json:"responseCodeDescription,omitempty"`
	TotalAuthorizedAmount   *AmountData                   `json:"totalAuthorizedAmount,omitempty"`
}

// NewAPIIncrementResponse constructs a new APIIncrementResponse instance.
func NewAPIIncrementResponse() *APIIncrementResponse {
	return &APIIncrementResponse{}
}
