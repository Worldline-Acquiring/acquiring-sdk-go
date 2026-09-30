// This file was automatically generated.

package domain

// PaginationResponse represents class PaginationResponse.
type PaginationResponse struct {
	LastIndex  *int32  `json:"lastIndex,omitempty"`
	SearchID   *string `json:"searchId,omitempty"`
	TotalCount *int32  `json:"totalCount,omitempty"`
}

// NewPaginationResponse constructs a new PaginationResponse instance.
func NewPaginationResponse() *PaginationResponse {
	return &PaginationResponse{}
}
