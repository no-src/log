package log

import (
	"errors"
	"testing"
	"time"

	"github.com/no-src/log/formatter"
	"github.com/no-src/log/level"
)

func TestMultiLogger(t *testing.T) {
	testCases := []struct {
		name      string
		formatter string
	}{
		{"TextFormatter", formatter.TextFormatter},
		{"JsonFormatter", formatter.JsonFormatter},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InitDefaultLogger(NewMultiLogger(NewConsoleLogger(level.DebugLevel).WithFormatter(formatter.New(tc.formatter))))
			defer Close()
			testLogs(t)
		})
	}
}

func TestMultiLogger_WithFormatter(t *testing.T) {
	testCases := []struct {
		name      string
		formatter string
	}{
		{"TextFormatter", formatter.TextFormatter},
		{"JsonFormatter", formatter.JsonFormatter},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// the WithFormatter is ineffective in the multiLogger
			InitDefaultLogger(NewMultiLogger(NewConsoleLogger(level.DebugLevel)).WithFormatter(formatter.New(tc.formatter)))
			defer Close()
			testLogs(t)
		})
	}
}

func TestMultiLogger_WithTimeFormat(t *testing.T) {
	testCases := []struct {
		name   string
		format string
	}{
		{"empty", ""},
		{"default", testTimeFormat},
		{"RFC3339", time.RFC3339},
		{"RFC3339Nano", time.RFC3339Nano},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// the WithTimeFormat is ineffective in the multiLogger
			InitDefaultLogger(NewMultiLogger(NewConsoleLogger(level.DebugLevel)).WithTimeFormat(tc.format))
			defer Close()
			testLogs(t)
		})
	}
}

func TestMultiLogger_WithError(t *testing.T) {
	InitDefaultLogger(NewMultiLogger(newErrorLogger(level.DebugLevel)))
	defer Close()
	testLogs(t)
}

type errorLogger struct {
	emptyLogger
}

func newErrorLogger(lvl level.Level) Logger {
	logger := &errorLogger{}
	return logger
}

func (l *errorLogger) Write(p []byte) (n int, err error) {
	return 0, errors.New("write to error logger")
}

func (l *errorLogger) Close() error {
	return errors.New("close the error logger")
}
