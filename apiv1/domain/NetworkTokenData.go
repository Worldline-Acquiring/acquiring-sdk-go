// This file was automatically generated.

package domain

// NetworkTokenData represents class NetworkTokenData
type NetworkTokenData struct {
	Cryptogram *string `json:"cryptogram,omitempty"`
	Eci        *string `json:"eci,omitempty"`
}

// NewNetworkTokenData constructs a new NetworkTokenData instance
func NewNetworkTokenData() *NetworkTokenData {
	return &NetworkTokenData{}
}
