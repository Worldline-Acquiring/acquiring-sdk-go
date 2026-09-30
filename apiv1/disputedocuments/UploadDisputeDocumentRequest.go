// This file was automatically generated.

package disputedocuments

import (
	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/domain"
)

// UploadDisputeDocumentRequest represents multipart/form-data parameters for Upload Dispute Document.
//
// Documentation can be found at https://docs.acquiring.worldline-solutions.com/api-reference#tag/Dispute-Documents/operation/uploadDisputeDocument.
type UploadDisputeDocumentRequest struct {
	File *domain.UploadableFile
}

// ToMultipartFormDataObject converts the multipart/form-data request to communication.MultipartFormDataObject.
func (request UploadDisputeDocumentRequest) ToMultipartFormDataObject() *communication.MultipartFormDataObject {
	multipart, _ := communication.NewMultipartFormDataObject()

	if request.File != nil {
		_ = multipart.AddFile("file", *request.File)
	}

	return multipart
}

// NewUploadDisputeDocumentRequest constructs a new UploadDisputeDocumentRequest instance.
func NewUploadDisputeDocumentRequest() *UploadDisputeDocumentRequest {
	return &UploadDisputeDocumentRequest{}
}
