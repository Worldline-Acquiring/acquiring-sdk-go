// This file was automatically generated.

package domain

// GeoCoordinates represents class GeoCoordinates
type GeoCoordinates struct {
	Latitude  *float64 `json:"latitude,omitempty"`
	Longitude *float64 `json:"longitude,omitempty"`
}

// NewGeoCoordinates constructs a new GeoCoordinates instance
func NewGeoCoordinates() *GeoCoordinates {
	return &GeoCoordinates{}
}
