package util

import (
	"strings"
	"time"
)

const timeFormat = "2006-01-02 15:04:05"

// CustomTime is used to implement a custom JSON deserialization.
type CustomTime struct {
	time.Time
}

// UnmarshalJSON is a custom JSON deserialization method used
// to parse the valued in timeFormat into a Time value.
func (ct *CustomTime) UnmarshalJSON(b []byte) (err error) {
	// trim the quotes
	s := strings.Trim(string(b), "\"")
	if s == "null" {
		ct.Time = time.Time{}
		return
	}

	ct.Time, err = time.Parse(timeFormat, s)
	return
}
