package content

import "time"

const (
	// DefaultLogTimeFormat the default log time format
	DefaultLogTimeFormat = "2006-01-02 15:04:05"
)

// Time the custom Time for log
type Time struct {
	time   time.Time
	format string
}

// NewTime convert time.Time to content.Time pointer with default format
func NewTime(time time.Time) *Time {
	return NewTimeWithFormat(time, DefaultLogTimeFormat)
}

// NewTimeWithFormat convert time.Time to content.Time pointer with custom format
func NewTimeWithFormat(time time.Time, format string) *Time {
	nt := Time{
		time:   time,
		format: format,
	}
	return &nt
}

// MarshalText implement interface encoding.TextMarshaler
func (t Time) MarshalText() (text []byte, err error) {
	return []byte(t.String()), nil
}

// Time convert to time.Time
func (t Time) Time() time.Time {
	return t.time
}

// String return a formatted time string
func (t Time) String() string {
	return t.Time().Format(t.format)
}
