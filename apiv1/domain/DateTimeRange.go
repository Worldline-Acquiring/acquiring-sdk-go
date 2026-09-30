// This file was automatically generated.

package domain

// DateTimeRange represents class DateTimeRange.
type DateTimeRange struct {
	Greater      *string `json:"greater,omitempty"`
	GreaterEqual *string `json:"greaterEqual,omitempty"`
	Lower        *string `json:"lower,omitempty"`
	LowerEqual   *string `json:"lowerEqual,omitempty"`
}

// NewDateTimeRange constructs a new DateTimeRange instance.
func NewDateTimeRange() *DateTimeRange {
	return &DateTimeRange{}
}
