package flights

import "strings"

// Day represents a day on the calendar from Spirits website
type Day struct {
	text String
}

// String is an alias for string to be able to use extension methods
type String string

// FlightAvailable checks a day to see if there's a flight available
func (d *Day) FlightAvailable() bool {
	return !strings.Contains(strings.ToUpper(string(d.text)), "NOT AVAILABLE")
}

// ConvertToDay converts a string to a day
func (s String) ConvertToDay() *Day {
	return &Day{text: s}
}

// ConvertToDays converts an array of strings to days
func ConvertToDays(s []*String) []*Day {
	days := []*Day{}

	for _, text := range s {
		days = append(days, text.ConvertToDay())
	}

	return days
}
