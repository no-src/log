package content

import (
	"io"
	"testing"
	"time"

	"github.com/no-src/log/level"
)

func TestNewContent(t *testing.T) {
	testCases := []struct {
		name       string
		err        error
		appendTime bool
		timeFormat string
	}{
		{"with nil error and no append time", nil, false, time.RFC3339},
		{"with error and no append time", io.EOF, false, time.RFC3339},
		{"with nil error and append time", nil, true, time.RFC3339},
		{"with error and append time", io.EOF, true, time.RFC3339},
		{"with empty time format", io.EOF, true, ""},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			c := NewContent(level.DebugLevel, tc.err, tc.appendTime, tc.timeFormat, tc.name)
			if c.AppendTime != tc.appendTime {
				t.Errorf("test NewContent failed, expect to get %v, but actual get %v", tc.appendTime, c.AppendTime)
				return
			}
			if tc.appendTime && c.Time == nil {
				t.Errorf("test NewContent failed, time can't be nil")
				return
			}

			if !tc.appendTime && c.Time != nil {
				t.Errorf("test NewContent failed, time should be nil")
				return
			}

			if tc.err != nil && c.Error == nil {
				t.Errorf("test NewContent failed, error can't be nil")
			}

			if tc.err == nil && c.Error != nil {
				t.Errorf("test NewContent failed, error should be nil")
			}
		})
	}
}
