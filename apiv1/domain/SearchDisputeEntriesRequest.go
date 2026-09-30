// This file was automatically generated.

package domain

// SearchDisputeEntriesRequest represents class SearchDisputeEntriesRequest.
type SearchDisputeEntriesRequest struct {
	DisputeID             *string                `json:"disputeId,omitempty"`
	EntryCategories       *[]string              `json:"entryCategories,omitempty"`
	EntryDateTime         *DateTimeRange         `json:"entryDateTime,omitempty"`
	EntryID               *string                `json:"entryId,omitempty"`
	EntryTypes            *[]string              `json:"entryTypes,omitempty"`
	IncludeDisputeSummary *bool                  `json:"includeDisputeSummary,omitempty"`
	MerchantScope         MerchantScopeInterface `json:"merchantScope,omitempty"`
	Pagination            *PaginationRequest     `json:"pagination,omitempty"`
	SortOrder             *string                `json:"sortOrder,omitempty"`
}

// NewSearchDisputeEntriesRequest constructs a new SearchDisputeEntriesRequest instance.
func NewSearchDisputeEntriesRequest() *SearchDisputeEntriesRequest {
	return &SearchDisputeEntriesRequest{}
}
