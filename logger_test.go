package log

import (
	"errors"
	"testing"

	"github.com/no-src/log/formatter"
	"github.com/no-src/log/level"
)

func testLogs(t *testing.T) {
	Debug("%s %s, test debug log", "hello", "world")
	Info("%s %s, test info log", "hello", "world")
	Warn("%s %s, test warn log", "hello", "world")
	Error(errors.New("log err"), "%s %s,test error log", "hello", "world")
	ErrorIf(errors.New("log err from ErrorIf"), "%s %s, test error log", "hello", "world")
	ErrorIf(nil, "%s %s, this error log will not be printed", "hello", "world")
	testSampleLogs()
	Log("%s %s, test log log", "hello", "world")
	Log("%s %s, test log log again", "hello", "world")
	DefaultLogger().Write([]byte(""))
	DefaultLogger().Write([]byte("hello logger"))
}

func TestDefaultLogger(t *testing.T) {
	defer Close()
	testLogs(t)
}

func TestConsoleLogger(t *testing.T) {
	testCases := []struct {
		name      string
		formatter string
	}{
		{"TextFormatter", formatter.TextFormatter},
		{"JsonFormatter", formatter.JsonFormatter},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InitDefaultLogger(NewConsoleLogger(level.DebugLevel).WithFormatter(formatter.New(tc.formatter)))
			defer Close()
			testLogs(t)
		})
	}
}

func TestEmptyLogger(t *testing.T) {
	testCases := []struct {
		name      string
		formatter string
	}{
		{"TextFormatter", formatter.TextFormatter},
		{"JsonFormatter", formatter.JsonFormatter},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InitDefaultLogger(NewEmptyLogger().WithFormatter(formatter.New(tc.formatter)))
			defer Close()
			testLogs(t)
		})
	}
}

func TestMinLogLevel(t *testing.T) {
	InitDefaultLogger(NewConsoleLogger(level.InfoLevel))
	defer Close()
	testLogs(t)
}

func TestNilLogger(t *testing.T) {
	InitDefaultLogger(nil)
	defer Close()
	testLogs(t)
}

func TestBaseLogger_Close(t *testing.T) {
	InitDefaultLogger(newMinLogger())
	// call baseLogger.Close
	defer Close()
	testLogs(t)
}

type minLogger struct {
	baseLogger
}

func newMinLogger() Logger {
	logger := &minLogger{}
	logger.init(logger, level.DebugLevel, true)
	return logger
}

func (l *minLogger) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (l *minLogger) WithFormatter(f formatter.Formatter) Logger {
	return l
}
