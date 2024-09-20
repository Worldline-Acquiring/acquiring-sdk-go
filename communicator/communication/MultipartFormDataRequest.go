package communication

// MultipartFormDataRequest represents a multipart/form-data request.
type MultipartFormDataRequest interface {

	// ToMultipartFormDataObject converts this multipart/form-data request into a MultipartFormDataObject.
	ToMultipartFormDataObject() *MultipartFormDataObject
}
