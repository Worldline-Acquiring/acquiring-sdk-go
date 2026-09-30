// This file was automatically generated.

package domain

// DisputeDateTimeData represents class DisputeDateTimeData.
type DisputeDateTimeData struct {
	ClosedDateTime            *string `json:"closedDateTime,omitempty"`
	LastStatusChangedDateTime *string `json:"lastStatusChangedDateTime,omitempty"`
	OpenedDateTime            *string `json:"openedDateTime,omitempty"`
	ResponseDueDate           *string `json:"responseDueDate,omitempty"`
}

// NewDisputeDateTimeData constructs a new DisputeDateTimeData instance.
func NewDisputeDateTimeData() *DisputeDateTimeData {
	return &DisputeDateTimeData{}
}
