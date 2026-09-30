// This file was automatically generated.

package domain

// PaginationRequest represents class PaginationRequest.
type PaginationRequest struct {
	FromIndex *int32  `json:"fromIndex,omitempty"`
	PageSize  *int32  `json:"pageSize,omitempty"`
	SearchID  *string `json:"searchId,omitempty"`
}

// NewPaginationRequest constructs a new PaginationRequest instance.
func NewPaginationRequest() *PaginationRequest {
	return &PaginationRequest{}
}
