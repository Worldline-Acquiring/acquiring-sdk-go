package acquiringsdk

import "time"

const dateLayout = "2006-01-02"

const dateTimeLayout = "2006-01-02T15:04:05.999Z07:00"

var altDateTimeLayouts = []string{
	"2006-01-02T15:04:05.999999999Z07:00",
	"2006-01-02T15:04Z07:00",
}

// ParseDate parses a string in ISO format into a time.Time instance
func ParseDate(date string) (time.Time, error) {
	return time.Parse(dateLayout, date)
}

// FormatDate formats a time.Time instance into an ISO date string
func FormatDate(date time.Time) string {
	return date.Format(dateLayout)
}

// ParseDateTime parses a string in ISO into a time.Time instance
func ParseDateTime(dateTime string) (time.Time, error) {
	for _, layout := range altDateTimeLayouts {
		result, err := time.Parse(layout, dateTime)
		if err == nil {
			return result, nil
		}
	}
	return time.Parse(dateTimeLayout, dateTime)
}

// FormatDateTime formats a time.Time instance into an ISO date-time string
func FormatDateTime(dateTime time.Time) string {
	return dateTime.Format(dateTimeLayout)
}
