// This file was automatically generated.

package domain

// APIReversalResponse represents class ApiReversalResponse
type APIReversalResponse struct {
	OperationID             *string                       `json:"operationId,omitempty"`
	Payment                 *APIPaymentSummaryForResponse `json:"payment,omitempty"`
	Responder               *string                       `json:"responder,omitempty"`
	ResponseCode            *string                       `json:"responseCode,omitempty"`
	ResponseCodeCategory    *string                       `json:"responseCodeCategory,omitempty"`
	ResponseCodeDescription *string                       `json:"responseCodeDescription,omitempty"`
	TotalAuthorizedAmount   *AmountData                   `json:"totalAuthorizedAmount,omitempty"`
}

// NewAPIReversalResponse constructs a new APIReversalResponse instance
func NewAPIReversalResponse() *APIReversalResponse {
	return &APIReversalResponse{}
}
