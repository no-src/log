package log

import (
	"errors"
	"testing"

	"github.com/no-src/log/formatter"
	"github.com/no-src/log/level"
)

func TestMultiLogger(t *testing.T) {
	testCases := []struct {
		name      string
		formatter formatter.Type
	}{
		{"TextFormatter", formatter.TextFormatter},
		{"JsonFormatter", formatter.JsonFormatter},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InitDefaultLogger(NewMultiLogger(NewConsoleLogger(level.DebugLevel).WithFormatter(formatter.New(tc.formatter))))
			defer Close()
			TestLogs(t)
		})
	}
}

func TestMultiLogger_WithFormatter(t *testing.T) {
	testCases := []struct {
		name      string
		formatter formatter.Type
	}{
		{"TextFormatter", formatter.TextFormatter},
		{"JsonFormatter", formatter.JsonFormatter},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// the WithFormatter is ineffective in the multiLogger
			InitDefaultLogger(NewMultiLogger(NewConsoleLogger(level.DebugLevel)).WithFormatter(formatter.New(tc.formatter)))
			defer Close()
			TestLogs(t)
		})
	}
}

func TestMultiLogger_WithError(t *testing.T) {
	InitDefaultLogger(NewMultiLogger(newErrorLogger(level.DebugLevel)))
	defer Close()
	TestLogs(t)
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
