package log

import (
	"io"

	"github.com/no-src/log/formatter"
	_ "github.com/no-src/log/formatter/json" // register json formatter
	_ "github.com/no-src/log/formatter/text" // register text formatter
)

// Logger define a universal log interface
type Logger interface {
	Writer
	Option

	Debug(format string, args ...any)
	Info(format string, args ...any)
	Warn(format string, args ...any)
	Error(err error, format string, args ...any)
}

// Writer implement write to log
type Writer interface {
	io.Writer

	// Log write log to output
	Log(format string, args ...any)
	// Close to close log and release dependencies
	Close() error
}

// Option the log options interface
type Option interface {
	// WithFormatter set the log formatter and return logger self
	WithFormatter(f formatter.Formatter) Logger
	// WithTimeFormat set the time format and return logger self
	WithTimeFormat(f string) Logger
}
