package log

import (
	"errors"
	"testing"

	"github.com/no-src/log/content"
	"github.com/no-src/log/formatter"
	"github.com/no-src/log/level"
)

func TestSampleLogger(t *testing.T) {
	testCases := []struct {
		name       string
		sampleRate float64
		formatter  string
	}{
		{"sample rate less than zero", -1, formatter.TextFormatter},
		{"sample rate equals zero", 0, formatter.TextFormatter},
		{"normal sample rate", 0.5, formatter.TextFormatter},
		{"sample rate equals one", 1, formatter.TextFormatter},
		{"sample rate greater than one", 2, formatter.TextFormatter},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InitDefaultLogger(NewDefaultSampleLogger(NewConsoleLogger(level.DebugLevel), tc.sampleRate).WithFormatter(formatter.New(tc.formatter)).WithTimeFormat(testTimeFormat))
			testLogs(t)
			Close()
		})
	}
}

func testSampleLogs() {
	DefaultSampleLogger().WithFormatter(formatter.Default()).WithTimeFormat(content.DefaultLogTimeFormat())
	DebugSample("[sample] %s %s, test debug log", "hello", "world")
	InfoSample("[sample] %s %s, test info log", "hello", "world")
	WarnSample("[sample] %s %s, test warn log", "hello", "world")
	ErrorSample(errors.New("log err"), "[sample] %s %s,test error log", "hello", "world")
	ErrorIfSample(errors.New("log err from ErrorIfSample"), "[sample] %s %s, test error log", "hello", "world")
	ErrorIfSample(nil, "[sample] %s %s, this error log will not be printed", "hello", "world")
}
