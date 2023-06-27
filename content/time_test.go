package content

import (
	"testing"
	"time"
)

func TestTime_MarshalText(t *testing.T) {
	testCases := []struct {
		name   string
		format string
	}{
		{"default", ""},
		{"RFC3339", time.RFC3339},
		{"RFC3339Nano", time.RFC3339Nano},
	}
	for _, tc := range testCases {
		now := time.Now()
		if len(tc.format) > 0 {
			InitDefaultLogTimeFormat(tc.format)
		} else {
			tc.format = DefaultLogTimeFormat()
		}

		expect := now.Format(tc.format)
		ti := NewTime(time.Now())
		data, err := ti.MarshalText()
		if err != nil {
			t.Errorf("Time.MarshalText error => %v", err)
			return
		}
		actual := string(data)
		if expect != actual {
			t.Errorf("test Time.MarshalText failed, expect to get %s, but actual get %s", expect, actual)
		}
	}

}
