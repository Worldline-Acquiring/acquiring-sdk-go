// This file was automatically generated.

package domain

// ThreeDSecure represents class ThreeDSecure
type ThreeDSecure struct {
	AuthenticationValue          *string `json:"authenticationValue,omitempty"`
	DirectoryServerTransactionID *string `json:"directoryServerTransactionId,omitempty"`
	Eci                          *string `json:"eci,omitempty"`
	ThreeDSecureType             *string `json:"threeDSecureType,omitempty"`
	Version                      *string `json:"version,omitempty"`
}

// NewThreeDSecure constructs a new ThreeDSecure instance
func NewThreeDSecure() *ThreeDSecure {
	return &ThreeDSecure{}
}
