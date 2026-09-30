// This file was automatically generated.

package domain

// SubmitEvidenceRequest represents class SubmitEvidenceRequest.
type SubmitEvidenceRequest struct {
	DocumentIds    *[]DisputeDocumentIDItem `json:"documentIds,omitempty"`
	Elaboration    *string                  `json:"elaboration,omitempty"`
	IncludeEntries *bool                    `json:"includeEntries,omitempty"`
	PartialAmount  *AmountData              `json:"partialAmount,omitempty"`
	UserID         *string                  `json:"userId,omitempty"`
}

// NewSubmitEvidenceRequest constructs a new SubmitEvidenceRequest instance.
func NewSubmitEvidenceRequest() *SubmitEvidenceRequest {
	return &SubmitEvidenceRequest{}
}
