// This file was automatically generated.

package domain

// APIPaymentErrorResponse represents class ApiPaymentErrorResponse
type APIPaymentErrorResponse struct {
	Detail   *string `json:"detail,omitempty"`
	Instance *string `json:"instance,omitempty"`
	Status   *int32  `json:"status,omitempty"`
	Title    *string `json:"title,omitempty"`
	Type     *string `json:"type,omitempty"`
}

// NewAPIPaymentErrorResponse constructs a new APIPaymentErrorResponse instance
func NewAPIPaymentErrorResponse() *APIPaymentErrorResponse {
	return &APIPaymentErrorResponse{}
}
