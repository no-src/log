package log

import (
	"os"

	"github.com/no-src/log/level"
)

type consoleLogger struct {
	baseLogger
}

// NewConsoleLogger get a console logger
func NewConsoleLogger(lvl level.Level) Logger {
	logger := &consoleLogger{}
	// init baseLogger
	logger.baseLogger.init(logger, lvl, true)
	return logger
}

func (l *consoleLogger) Write(p []byte) (n int, err error) {
	return os.Stdout.Write(p)
}
