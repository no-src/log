package log

import (
	"github.com/no-src/log/formatter"
)

type emptyLogger struct {
}

// NewEmptyLogger get an empty logger, there is nothing to do
func NewEmptyLogger() Logger {
	logger := &emptyLogger{}
	return logger
}

func (l *emptyLogger) Debug(format string, args ...any) {

}

func (l *emptyLogger) Info(format string, args ...any) {

}

func (l *emptyLogger) Warn(format string, args ...any) {

}

func (l *emptyLogger) Error(err error, format string, args ...any) {

}

func (l *emptyLogger) Log(format string, args ...any) {

}

func (l *emptyLogger) Close() error {
	return nil
}

func (l *emptyLogger) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (l *emptyLogger) WithFormatter(f formatter.Formatter) Logger {
	return l
}

func (l *emptyLogger) WithTimeFormat(f string) Logger {
	return l
}
