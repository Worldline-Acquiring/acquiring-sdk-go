// This file was automatically generated.

package domain

// APIActionResponse represents class ApiActionResponse
type APIActionResponse struct {
	OperationID             *string                       `json:"operationId,omitempty"`
	Payment                 *APIPaymentSummaryForResponse `json:"payment,omitempty"`
	Responder               *string                       `json:"responder,omitempty"`
	ResponseCode            *string                       `json:"responseCode,omitempty"`
	ResponseCodeCategory    *string                       `json:"responseCodeCategory,omitempty"`
	ResponseCodeDescription *string                       `json:"responseCodeDescription,omitempty"`
}

// NewAPIActionResponse constructs a new APIActionResponse instance
func NewAPIActionResponse() *APIActionResponse {
	return &APIActionResponse{}
}
