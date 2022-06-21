package content

import "time"

// Time the custom Time for log
type Time time.Time

// MarshalText implement interface encoding.TextMarshaler
func (t Time) MarshalText() (text []byte, err error) {
	return []byte(time.Time(t).Format("2006-01-02 15:04:05")), nil
}