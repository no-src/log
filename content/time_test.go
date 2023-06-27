package content

import (
	"sync"
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
		ti := NewTime(now)
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

func TestInitDefaultLogTimeFormat_Concurrency(t *testing.T) {
	c := 10
	wg := sync.WaitGroup{}
	wg.Add(c * 2)
	for i := 0; i < c; i++ {
		go func() {
			InitDefaultLogTimeFormat(time.RFC3339)
			wg.Done()
		}()

		go func() {
			NewTime(time.Now())
			wg.Done()
		}()
	}
	wg.Wait()
}
