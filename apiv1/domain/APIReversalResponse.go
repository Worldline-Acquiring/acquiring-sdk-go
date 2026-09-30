// This file was automatically generated.

package domain

// APIReversalResponse represents class ApiReversalResponse.
type APIReversalResponse struct {
	APIActionResponse
	TotalAuthorizedAmount *AmountData `json:"totalAuthorizedAmount,omitempty"`
}

// NewAPIReversalResponse constructs a new APIReversalResponse instance.
func NewAPIReversalResponse() *APIReversalResponse {
	return &APIReversalResponse{}
}
