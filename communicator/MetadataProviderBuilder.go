package communicator

import (
	"errors"
	"strings"

	"github.com/Worldline-Acquiring/acquiring-sdk-go/communicator/communication"
	"github.com/Worldline-Acquiring/acquiring-sdk-go/domain"
)

// MetadataProviderBuilder represents a builder for a MetadataProvider object.
type MetadataProviderBuilder struct {
	integrator               string
	ShoppingCartExtension    *domain.ShoppingCartExtension
	AdditionalRequestHeaders []communication.Header
}

// Build creates a fully initialized MetadataProvider
func (m MetadataProviderBuilder) Build() (*MetadataProvider, error) {
	return newMetadataProvider(m.integrator, m.ShoppingCartExtension, m.AdditionalRequestHeaders)
}

// NewMetadataProviderBuilder creates a MetadataProviderBuilder with the given Integrator
func NewMetadataProviderBuilder(integrator string) (*MetadataProviderBuilder, error) {
	if strings.TrimSpace(integrator) == "" {
		return nil, errors.New("integrator is required")
	}

	return &MetadataProviderBuilder{integrator, nil, nil}, nil
}
