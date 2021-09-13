package log

import "fmt"

type consoleLogger struct {
	baseLogger
	level Level
}

// NewConsoleLogger get a console logger
func NewConsoleLogger(level Level) Logger {
	logger := &consoleLogger{
		level: level,
	}
	// use current instance Writer to override baseLogger's Writer
	logger.baseLogger.Writer = logger
	return logger
}

// Log write a format log to console
func (l *consoleLogger) Log(format string, args ...interface{}) {
	format = fmt.Sprintf(format, args...)
	format = appendRowTerminator(format)
	fmt.Print(format)
}
