// This file was automatically generated.

package domain

// SearchDisputesRequest represents class SearchDisputesRequest.
type SearchDisputesRequest struct {
	AcquirerDisputeReference  *string                `json:"acquirerDisputeReference,omitempty"`
	AcquirerReferenceNumber   *string                `json:"acquirerReferenceNumber,omitempty"`
	ClosedDateTime            *DateTimeRange         `json:"closedDateTime,omitempty"`
	DisputeID                 *string                `json:"disputeId,omitempty"`
	DisputeStages             *[]string              `json:"disputeStages,omitempty"`
	DisputeStatusCategories   *[]string              `json:"disputeStatusCategories,omitempty"`
	IsOpen                    *bool                  `json:"isOpen,omitempty"`
	LastStatusChangedDateTime *DateTimeRange         `json:"lastStatusChangedDateTime,omitempty"`
	MerchantReference         *string                `json:"merchantReference,omitempty"`
	MerchantScope             MerchantScopeInterface `json:"merchantScope,omitempty"`
	OpenedDateTime            *DateTimeRange         `json:"openedDateTime,omitempty"`
	Pagination                *PaginationRequest     `json:"pagination,omitempty"`
	PaymentID                 *string                `json:"paymentId,omitempty"`
	ResponseDueDate           *DateRange             `json:"responseDueDate,omitempty"`
	Schemes                   *[]string              `json:"schemes,omitempty"`
	SortBy                    *string                `json:"sortBy,omitempty"`
	SortOrder                 *string                `json:"sortOrder,omitempty"`
	UnifiedCategories         *[]string              `json:"unifiedCategories,omitempty"`
}

// NewSearchDisputesRequest constructs a new SearchDisputesRequest instance.
func NewSearchDisputesRequest() *SearchDisputesRequest {
	return &SearchDisputesRequest{}
}
