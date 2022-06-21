package log

import (
	"errors"
	"testing"

	"github.com/no-src/log/level"
)

func TestMultiLogger(t *testing.T) {
	InitDefaultLogger(NewMultiLogger(NewConsoleLogger(level.DebugLevel)))
	defer Close()
	TestLogs(t)
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
