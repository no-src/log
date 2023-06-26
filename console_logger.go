package log

import (
	"bufio"
	"io"
	"os"

	"github.com/no-src/log/formatter"
	"github.com/no-src/log/level"
)

type consoleLogger struct {
	baseLogger

	w          io.Writer
	bw         *bufio.Writer
	withBuffer bool
}

// NewConsoleLogger get a console logger
func NewConsoleLogger(lvl level.Level) Logger {
	return newConsoleLoggerWithBuffer(lvl, false)
}

// newConsoleLoggerWithBuffer get a console logger with buffer
func newConsoleLoggerWithBuffer(lvl level.Level, withBuffer bool) Logger {
	logger := &consoleLogger{
		w:          os.Stdout,
		bw:         bufio.NewWriter(os.Stdout),
		withBuffer: withBuffer,
	}

	// init baseLogger
	logger.baseLogger.init(logger, lvl, true)
	return logger
}

func (l *consoleLogger) Write(p []byte) (n int, err error) {
	if l.withBuffer {
		return l.bw.Write(p)
	}
	return l.w.Write(p)
}

func (l *consoleLogger) Close() error {
	if l.withBuffer {
		return l.bw.Flush()
	}
	return nil
}

func (l *consoleLogger) WithFormatter(f formatter.Formatter) Logger {
	l.setFormatter(f)
	return l
}

func (l *consoleLogger) WithTimeFormat(f string) Logger {
	l.setTimeFormat(f)
	return l
}
