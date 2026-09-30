// This file was automatically generated.

package domain

// DisputeDocument represents class DisputeDocument.
type DisputeDocument struct {
	DocumentID *string `json:"documentId,omitempty"`
	FileName   *string `json:"fileName,omitempty"`
	MimeType   *string `json:"mimeType,omitempty"`
}

// NewDisputeDocument constructs a new DisputeDocument instance.
func NewDisputeDocument() *DisputeDocument {
	return &DisputeDocument{}
}
