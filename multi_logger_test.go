package log

import (
	"errors"
	"testing"

	"github.com/no-src/log/formatter"
	"github.com/no-src/log/level"
)

func TestMultiLogger(t *testing.T) {
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
			fLogger, err := NewFileLogger(level.DebugLevel, "./multi_logs", "ns"+tc.formatter)
			if err != nil {
				t.Fatal(err)
			}
			InitDefaultLogger(NewMultiLogger(NewConsoleLogger(level.DebugLevel), fLogger).WithFormatter(formatter.New(tc.formatter)).WithTimeFormat(tc.timeFormat))
			defer Close()

			if tc.concurrency {
				testLogsConcurrency(t, "TestMultiLogger")
			} else {
				testLogs(t)
			}
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
		{"customized", testTimeFormat},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
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
