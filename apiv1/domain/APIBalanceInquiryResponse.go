// This file was automatically generated.

package domain

// APIBalanceInquiryResponse represents class ApiBalanceInquiryResponse
type APIBalanceInquiryResponse struct {
	AuthorizationCode       *string                    `json:"authorizationCode,omitempty"`
	AvailableAmount         *AmountData                `json:"availableAmount,omitempty"`
	OperationID             *string                    `json:"operationId,omitempty"`
	References              *APIReferencesForResponses `json:"references,omitempty"`
	Responder               *string                    `json:"responder,omitempty"`
	ResponseCode            *string                    `json:"responseCode,omitempty"`
	ResponseCodeCategory    *string                    `json:"responseCodeCategory,omitempty"`
	ResponseCodeDescription *string                    `json:"responseCodeDescription,omitempty"`
}

// NewAPIBalanceInquiryResponse constructs a new APIBalanceInquiryResponse instance
func NewAPIBalanceInquiryResponse() *APIBalanceInquiryResponse {
	return &APIBalanceInquiryResponse{}
}
