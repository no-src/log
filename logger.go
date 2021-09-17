package log

import (
	"fmt"
	"strings"
)

// Logger define a universal log interface
type Logger interface {
	Debug(format string, args ...interface{})
	Info(format string, args ...interface{})
	Warn(format string, args ...interface{})
	Error(err error, format string, args ...interface{})
	Writer
}

// Writer implement write to log
type Writer interface {
	// Log write log to output
	Log(format string, args ...interface{})
	// Close to close log and release dependencies
	Close() error
}

var loggerFormat = "[%s] %s"          // [level] content
var errorLoggerFormat = "[%s] %s. %s" // [level] content. error
var defaultTerminator = "\n"

func buildLog(level Level, format string) string {
	format = fmt.Sprintf(loggerFormat, level.String(), format)
	return format
}

func buildErrorLog(level Level, err error, format string) string {
	format = fmt.Sprintf(errorLoggerFormat, level.String(), format, err)
	return format
}

func appendRowTerminator(format string) string {
	if !strings.HasSuffix(format, defaultTerminator) {
		format = format + defaultTerminator
	}
	return format
}

func checkLogLevel(level Level, currentLevel Level) bool {
	return currentLevel >= level
}
