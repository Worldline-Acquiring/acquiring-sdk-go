package communicator

import (
	"errors"
	"strings"
)

// APIResource represents the base type of all Worldline Acquiring platform API resources.
type APIResource struct {
	parentResource *APIResource
	communicator   *Communicator
	pathContext    map[string]string
}

// Communicator returns the Communicator used by the resource
func (ar APIResource) Communicator() *Communicator {
	return ar.communicator
}

// InstantiateURIWithContext instantiates the given URI with the path context
func (ar APIResource) InstantiateURIWithContext(uri string, pathContext map[string]string) (string, error) {
	return ar.InstantiateURI(replaceAll(uri, pathContext))
}

// InstantiateURI instantiates the given uri with the path context of the resource
func (ar APIResource) InstantiateURI(uri string) (string, error) {
	uri = replaceAll(uri, ar.pathContext)

	if ar.parentResource != nil {
		parentURI, err := ar.parentResource.InstantiateURI(uri)
		if err != nil {
			return parentURI, err
		}

		uri = parentURI
	}

	return uri, nil
}

func replaceAll(uri string, pathContext map[string]string) string {
	if pathContext != nil {
		for key, value := range pathContext {
			uri = strings.Replace(uri, "{"+key+"}", value, -1)
		}
	}

	return uri
}

// NewAPIResourceWithParent creates an APIResource with the given parent and pathContext
func NewAPIResourceWithParent(parent *APIResource, pathContext map[string]string) (*APIResource, error) {
	if parent == nil {
		return nil, errors.New("parent is required")
	}

	return &APIResource{parent, parent.communicator, pathContext}, nil
}

// NewAPIResource creates an APIResource with the given communicator and pathContext
func NewAPIResource(communicator *Communicator, pathContext map[string]string) (*APIResource, error) {
	if communicator == nil {
		return nil, errors.New("communicator is required")
	}
	return &APIResource{nil, communicator, pathContext}, nil
}
