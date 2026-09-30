// This file was automatically generated.

package domain

// DateRange represents class DateRange.
type DateRange struct {
	GreaterEqual *string `json:"greaterEqual,omitempty"`
	LowerEqual   *string `json:"lowerEqual,omitempty"`
}

// NewDateRange constructs a new DateRange instance.
func NewDateRange() *DateRange {
	return &DateRange{}
}
