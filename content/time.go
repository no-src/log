package content

import (
	"sync"
	"time"
)

var (
	defaultLogTimeFormat = "2006-01-02 15:04:05"
	mu                   sync.RWMutex
)

// Time the custom Time for log
type Time struct {
	time   time.Time
	format string
}

// InitDefaultLogTimeFormat init the global default log time format
func InitDefaultLogTimeFormat(f string) {
	if len(f) > 0 {
		mu.Lock()
		defaultLogTimeFormat = f
		mu.Unlock()
	}
}

// DefaultLogTimeFormat return the default log time format
func DefaultLogTimeFormat() string {
	mu.RLock()
	defer mu.RUnlock()
	return defaultLogTimeFormat
}

// NewTime convert time.Time to content.Time pointer with default format
func NewTime(time time.Time) *Time {
	return NewTimeWithFormat(time, DefaultLogTimeFormat())
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
