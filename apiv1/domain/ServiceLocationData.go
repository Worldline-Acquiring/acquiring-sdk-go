// This file was automatically generated.

package domain

// ServiceLocationData represents class ServiceLocationData
type ServiceLocationData struct {
	Address        *ServiceLocationAddress `json:"address,omitempty"`
	GeoCoordinates *GeoCoordinates         `json:"geoCoordinates,omitempty"`
}

// NewServiceLocationData constructs a new ServiceLocationData instance
func NewServiceLocationData() *ServiceLocationData {
	return &ServiceLocationData{}
}
