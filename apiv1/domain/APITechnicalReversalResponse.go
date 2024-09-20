// This file was automatically generated.

package domain

// APITechnicalReversalResponse represents class ApiTechnicalReversalResponse
type APITechnicalReversalResponse struct {
	OperationID             *string `json:"operationId,omitempty"`
	Responder               *string `json:"responder,omitempty"`
	ResponseCode            *string `json:"responseCode,omitempty"`
	ResponseCodeCategory    *string `json:"responseCodeCategory,omitempty"`
	ResponseCodeDescription *string `json:"responseCodeDescription,omitempty"`
}

// NewAPITechnicalReversalResponse constructs a new APITechnicalReversalResponse instance
func NewAPITechnicalReversalResponse() *APITechnicalReversalResponse {
	return &APITechnicalReversalResponse{}
}
