package log

import (
	"errors"
	"testing"
)

func TestMultiLogger(t *testing.T) {
	InitDefaultLogger(NewMultiLogger(NewConsoleLogger(DebugLevel)))
	defer Close()
	TestLogs(t)
}

func TestMultiLogger_WithError(t *testing.T) {
	InitDefaultLogger(NewMultiLogger(newErrorLogger(DebugLevel)))
	defer Close()
	TestLogs(t)
}

type errorLogger struct {
	emptyLogger
}

func newErrorLogger(level Level) Logger {
	logger := &errorLogger{}
	return logger
}

func (l *errorLogger) Write(p []byte) (n int, err error) {
	return 0, errors.New("write to error logger")
}

func (l *errorLogger) Close() error {
	return errors.New("close the error logger")
}
