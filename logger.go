package log

import (
	"io"

	_ "github.com/no-src/log/formatter/text" // register text formatter
)

// Logger define a universal log interface
type Logger interface {
	Writer

	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(err error, format string, args ...interface{})
}

// Writer implement write to log
type Writer interface {
	io.Writer

	// Log write log to output
	Log(format string, args ...interface{})
	// Close to close log and release dependencies
	Close() error
}
