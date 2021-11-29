package log

import (
	"fmt"
	"os"
)

type consoleLogger struct {
	baseLogger
}

// NewConsoleLogger get a console logger
func NewConsoleLogger(level Level) Logger {
	logger := &consoleLogger{}
	// init baseLogger
	logger.baseLogger.init(logger, level)
	return logger
}

// Log write a format log to console
func (l *consoleLogger) Log(format string, args ...interface{}) {
	format = fmt.Sprintf(format, args...)
	format = l.builder.AppendRowTerminator(format)
	fmt.Print(format)
}

func (l *consoleLogger) Close() error {
	return nil
}

func (l *consoleLogger) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}
