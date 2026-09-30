// This file was automatically generated.

package domain

// APIIncrementResponse represents class ApiIncrementResponse.
type APIIncrementResponse struct {
	APIActionResponse
	AuthorizationCode     *string     `json:"authorizationCode,omitempty"`
	TotalAuthorizedAmount *AmountData `json:"totalAuthorizedAmount,omitempty"`
}

// NewAPIIncrementResponse constructs a new APIIncrementResponse instance.
func NewAPIIncrementResponse() *APIIncrementResponse {
	return &APIIncrementResponse{}
}
