package log

import (
	"testing"

	"github.com/no-src/log/formatter"
	"github.com/no-src/log/level"
)

func TestConsoleLogger(t *testing.T) {
	testCases := []struct {
		name        string
		formatter   string
		concurrency bool
		timeFormat  string
	}{
		{"TextFormatter", formatter.TextFormatter, false, testTimeFormat},
		{"JsonFormatter", formatter.JsonFormatter, false, testTimeFormat},
		{"TextFormatter Concurrency", formatter.TextFormatter, true, ""},
		{"JsonFormatter Concurrency", formatter.JsonFormatter, true, ""},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InitDefaultLogger(NewConsoleLogger(level.DebugLevel).WithFormatter(formatter.New(tc.formatter)).WithTimeFormat(tc.timeFormat))
			defer Close()
			if tc.concurrency {
				testLogsConcurrency(t, "TestConsoleLogger")
			} else {
				testLogs(t)
			}
		})
	}
}

func TestConsoleLoggerWithBuffer(t *testing.T) {
	testCases := []struct {
		name      string
		formatter string
	}{
		{"TextFormatter", formatter.TextFormatter},
		{"JsonFormatter", formatter.JsonFormatter},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InitDefaultLogger(newConsoleLoggerWithBuffer(level.DebugLevel, true).WithFormatter(formatter.New(tc.formatter)).WithTimeFormat(testTimeFormat))
			defer Close()
			testLogs(t)
		})
	}
}
