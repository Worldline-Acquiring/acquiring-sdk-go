// This file was automatically generated.

package domain

// APIActionResponseForRefund represents class ApiActionResponseForRefund
type APIActionResponseForRefund struct {
	OperationID             *string                      `json:"operationId,omitempty"`
	Refund                  *APIRefundSummaryForResponse `json:"refund,omitempty"`
	Responder               *string                      `json:"responder,omitempty"`
	ResponseCode            *string                      `json:"responseCode,omitempty"`
	ResponseCodeCategory    *string                      `json:"responseCodeCategory,omitempty"`
	ResponseCodeDescription *string                      `json:"responseCodeDescription,omitempty"`
}

// NewAPIActionResponseForRefund constructs a new APIActionResponseForRefund instance
func NewAPIActionResponseForRefund() *APIActionResponseForRefund {
	return &APIActionResponseForRefund{}
}
