package content

import "time"

const (
	// LogTimeFormat the default log time format
	LogTimeFormat = "2006-01-02 15:04:05"
)

// Time the custom Time for log
type Time time.Time

// NewTime convert time.Time to content.Time pointer
func NewTime(t time.Time) *Time {
	nt := Time(t)
	return &nt
}

// MarshalText implement interface encoding.TextMarshaler
func (t Time) MarshalText() (text []byte, err error) {
	return []byte(time.Time(t).Format(LogTimeFormat)), nil
}

// Time convert to time.Time
func (t Time) Time() time.Time {
	return time.Time(t)
}
