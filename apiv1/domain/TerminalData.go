// This file was automatically generated.

package domain

// TerminalData represents class TerminalData.
type TerminalData struct {
	AllowSingleTap                   *bool     `json:"allowSingleTap,omitempty"`
	CardReadingCapabilities          *[]string `json:"cardReadingCapabilities,omitempty"`
	CardholderActivatedTerminalLevel *string   `json:"cardholderActivatedTerminalLevel,omitempty"`
	IsAttendedTerminal               *bool     `json:"isAttendedTerminal,omitempty"`
	IsOfflineApproved                *bool     `json:"isOfflineApproved,omitempty"`
	OfflineAuthorizationResponseCode *string   `json:"offlineAuthorizationResponseCode,omitempty"`
	PinEntryCapability               *string   `json:"pinEntryCapability,omitempty"`
	TerminalID                       *string   `json:"terminalId,omitempty"`
	TerminalLocation                 *string   `json:"terminalLocation,omitempty"`
}

// NewTerminalData constructs a new TerminalData instance.
func NewTerminalData() *TerminalData {
	return &TerminalData{}
}
