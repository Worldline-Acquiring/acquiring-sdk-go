// This file was automatically generated.

package domain

// ServiceLocationAddress represents class ServiceLocationAddress.
type ServiceLocationAddress struct {
	City                   *string `json:"city,omitempty"`
	CountryCode            *string `json:"countryCode,omitempty"`
	CountrySubdivisionCode *string `json:"countrySubdivisionCode,omitempty"`
	PostalCode             *string `json:"postalCode,omitempty"`
}

// NewServiceLocationAddress constructs a new ServiceLocationAddress instance.
func NewServiceLocationAddress() *ServiceLocationAddress {
	return &ServiceLocationAddress{}
}
