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
	Log(format string, args ...interface{})
}

var loggerFormat = "[%s] %s"          // [level] content
var errorLoggerFormat = "[%s] %s. %s" // [level] content. error
var defaultTerminator = "\n"

func buildLog(level Level, format string) string {
	format = fmt.Sprintf(loggerFormat, level.String(), format)
	return format
}

func buildErrorLog(level Level, err error, format string) string {
	format = fmt.Sprintf(errorLoggerFormat, level.String(), err, format)
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
