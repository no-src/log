package log

import (
	"bufio"
	"os"

	"github.com/no-src/log/formatter"
	"github.com/no-src/log/level"
)

type consoleLogger struct {
	baseLogger
	w *bufio.Writer
}

// NewConsoleLogger get a console logger
func NewConsoleLogger(lvl level.Level) Logger {
	logger := &consoleLogger{
		w: bufio.NewWriter(os.Stdout),
	}
	// init baseLogger
	logger.baseLogger.init(logger, lvl, true)
	return logger
}

func (l *consoleLogger) Write(p []byte) (n int, err error) {
	return l.w.Write(p)
}

func (l *consoleLogger) Close() error {
	return l.w.Flush()
}

func (l *consoleLogger) WithFormatter(f formatter.Formatter) Logger {
	if f != nil {
		l.f = f
	}
	return l
}
