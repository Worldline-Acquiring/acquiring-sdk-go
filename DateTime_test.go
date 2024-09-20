package acquiringsdk

import (
	"fmt"
	"testing"
	"time"
)

func TestParseDate(t *testing.T) {
	parameters := []struct {
		input    string
		expected time.Time
	}{
		{"2023-09-20", time.Date(2023, 9, 20, 0, 0, 0, 0, time.UTC)},
		{"2024-02-29", time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC)},
	}
	for i := range parameters {
		input := parameters[i].input
		expected := parameters[i].expected
		t.Run(fmt.Sprintf("applied to '%s'", input), func(t *testing.T) {
			actual, err := ParseDate(input)
			if err != nil {
				t.Fatalf("TestParseDate: %v", err)
			}
			if !areEqual(actual, expected) {
				t.Errorf("expected: '%s', actual: '%s'", expected, actual)
			}
		})
	}
}

func TestFormatDate(t *testing.T) {
	parameters := []struct {
		input    time.Time
		expected string
	}{
		{time.Date(2023, 9, 20, 0, 0, 0, 0, time.UTC), "2023-09-20"},
		{time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC), "2024-02-29"},
		{time.Date(2024, 2, 29, 0, 0, 0, 0, time.FixedZone("", 3600)), "2024-02-29"},
	}
	for i := range parameters {
		input := parameters[i].input
		expected := parameters[i].expected
		t.Run(fmt.Sprintf("applied to '%s'", input), func(t *testing.T) {
			actual := FormatDate(input)
			if actual != expected {
				t.Errorf("expected: '%s', actual: '%s'", expected, actual)
			}
		})
	}
}

func TestFormatAndParseDate(t *testing.T) {
	parameters := []time.Time{
		time.Date(2023, 9, 20, 0, 0, 0, 0, time.UTC),
		time.Date(2024, 2, 29, 0, 0, 0, 0, time.UTC),
	}
	for i := range parameters {
		input := parameters[i]
		t.Run(fmt.Sprintf("applied to '%s'", input), func(t *testing.T) {
			formatted := FormatDate(input)
			actual, err := ParseDate(formatted)
			if err != nil {
				t.Fatalf("TestFormatAndParseDate: %v", err)
			}
			if !areEqual(actual, input) {
				t.Errorf("expected: '%s', actual: '%s'", input, actual)
			}
		})
	}
}

func TestParseDateTime(t *testing.T) {
	parameters := []struct {
		input    string
		expected time.Time
	}{
		{"2023-10-10T08:00+02:00", time.Date(2023, 10, 10, 8, 0, 0, 0, time.FixedZone("", 7200))},
		{"2023-10-10T08:00Z", time.Date(2023, 10, 10, 8, 0, 0, 0, time.UTC)},
		{"2020-01-01T12:00:00Z", time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC)},
		{"2023-10-10T08:00:01.234+02:00", time.Date(2023, 10, 10, 8, 0, 1, 234000000, time.FixedZone("", 7200))},
		{"2023-10-10T08:00:01.234Z", time.Date(2023, 10, 10, 8, 0, 1, 234000000, time.UTC)},
		{"2020-01-01T12:00:00.123456Z", time.Date(2020, 1, 1, 12, 0, 0, 123456000, time.UTC)},
	}
	for i := range parameters {
		input := parameters[i].input
		expected := parameters[i].expected
		t.Run(fmt.Sprintf("applied to '%s'", input), func(t *testing.T) {
			actual, err := ParseDateTime(input)
			if err != nil {
				t.Fatalf("TestParseDateTime: %v", err)
			}
			if !areEqual(actual, expected) {
				t.Errorf("expected: '%s', actual: '%s'", expected, actual)
			}
		})
	}
}

func TestFormatDateTime(t *testing.T) {
	parameters := []struct {
		input    time.Time
		expected string
	}{
		{time.Date(2023, 10, 10, 8, 0, 0, 0, time.FixedZone("", 7200)), "2023-10-10T08:00:00+02:00"},
		{time.Date(2023, 10, 10, 8, 0, 0, 0, time.UTC), "2023-10-10T08:00:00Z"},
		{time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC), "2020-01-01T12:00:00Z"},
		{time.Date(2023, 10, 10, 8, 0, 1, 234000000, time.FixedZone("", 7200)), "2023-10-10T08:00:01.234+02:00"},
		{time.Date(2023, 10, 10, 8, 0, 1, 234000000, time.UTC), "2023-10-10T08:00:01.234Z"},
		{time.Date(2020, 1, 1, 12, 0, 0, 123456000, time.UTC), "2020-01-01T12:00:00.123Z"},
	}
	for i := range parameters {
		input := parameters[i].input
		expected := parameters[i].expected
		t.Run(fmt.Sprintf("applied to '%s'", input), func(t *testing.T) {
			actual := FormatDateTime(input)
			if actual != expected {
				t.Errorf("expected: '%s', actual: '%s'", expected, actual)
			}
		})
	}
}

func TestFormatAndParseDateTime(t *testing.T) {
	parameters := []time.Time{
		time.Date(2023, 10, 10, 8, 0, 0, 0, time.FixedZone("", 7200)),
		time.Date(2023, 10, 10, 8, 9, 9, 9, time.UTC),
		time.Date(2020, 1, 1, 12, 0, 0, 0, time.UTC),
		time.Date(2023, 10, 10, 8, 0, 1, 234000000, time.FixedZone("", 7200)),
		time.Date(2023, 10, 10, 8, 0, 1, 234000000, time.UTC),
		time.Date(2020, 1, 1, 12, 0, 0, 123456000, time.UTC),
	}
	for i := range parameters {
		input := parameters[i]
		t.Run(fmt.Sprintf("applied to '%s'", input), func(t *testing.T) {
			formatted := FormatDateTime(input)
			actual, err := ParseDateTime(formatted)
			// truncate the input nanos to millis
			nanos := input.Nanosecond() / 1000000 * 1000000
			expected := time.Date(input.Year(), input.Month(), input.Day(), input.Hour(), input.Minute(), input.Second(), nanos, input.Location())
			if err != nil {
				t.Fatalf("TestFormatAndParseDateTime: %v", err)
			}
			if !areEqual(actual, expected) {
				t.Errorf("expected: '%s', actual: '%s'", input, actual)
			}
		})
	}
}

func areEqual(a, b time.Time) bool {
	if a.UnixNano() != b.UnixNano() {
		return false
	}
	if a.Year() != b.Year() || a.Month() != b.Month() || a.Day() != b.Day() || a.Hour() != b.Hour() || a.Minute() != b.Minute() || a.Nanosecond() != b.Nanosecond() {
		return false
	}
	_, offsetA := a.Zone()
	_, offsetB := b.Zone()
	return offsetA == offsetB
}
