package log

import (
	"errors"
	"testing"

	"github.com/no-src/log/level"
)

func TestSampleLogger(t *testing.T) {
	testCases := []struct {
		name       string
		sampleRate float64
	}{
		{"sample rate less than zero", -1},
		{"sample rate equals zero", 0},
		{"normal sample rate", 0.5},
		{"sample rate equals one", 1},
		{"sample rate greater than one", 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InitDefaultLogger(NewDefaultSampleLogger(NewConsoleLogger(level.DebugLevel), tc.sampleRate))
			TestLogs(t)
			Close()
		})
	}
}

func testSampleLogs() {
	DebugSample("[sample] %s, test debug log", "hello")
	InfoSample("[sample] %s, test info log", "hello")
	WarnSample("[sample] %s, test warn log", "hello")
	ErrorSample(errors.New("log err"), "[sample] %s,test error log", "hello")
	ErrorIfSample(errors.New("log err from ErrorIfSample"), "[sample] %s, test error log", "hello")
	ErrorIfSample(nil, "[sample] %s, this error log will not be printed", "hello")
}
