// This file was automatically generated.

package domain

// SearchDisputesResponse represents class SearchDisputesResponse.
type SearchDisputesResponse struct {
	Disputes   *[]DisputeCase      `json:"disputes,omitempty"`
	Pagination *PaginationResponse `json:"pagination,omitempty"`
	RequestID  *string             `json:"requestId,omitempty"`
}

// NewSearchDisputesResponse constructs a new SearchDisputesResponse instance.
func NewSearchDisputesResponse() *SearchDisputesResponse {
	return &SearchDisputesResponse{}
}
