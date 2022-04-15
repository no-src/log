package log

import (
	"fmt"
	"strings"
)

// baseLogger Implement basic logger operation
type baseLogger struct {
	Writer
	builder

	level Level // min log level
}

func (l *baseLogger) Debug(format string, args ...interface{}) {
	if checkLogLevel(l.level, DebugLevel) {
		format = l.builder.BuildLog(DebugLevel, format)
		l.Writer.Log(format, args...)
	}
}

func (l *baseLogger) Info(format string, args ...interface{}) {
	if checkLogLevel(l.level, InfoLevel) {
		format = l.builder.BuildLog(InfoLevel, format)
		l.Writer.Log(format, args...)
	}
}

func (l *baseLogger) Warn(format string, args ...interface{}) {
	if checkLogLevel(l.level, WarnLevel) {
		format = l.builder.BuildLog(WarnLevel, format)
		l.Writer.Log(format, args...)
	}
}

func (l *baseLogger) Error(err error, format string, args ...interface{}) {
	if checkLogLevel(l.level, ErrorLevel) {
		format = l.builder.BuildErrorLog(ErrorLevel, err, format)
		l.Writer.Log(format, args...)
	}
}

var loggerFormat = "[%s] %s"                  // [level] content
var errorLoggerFormat = loggerFormat + ". %s" // [level] content. error
var defaultTerminator = "\n"

func (l *baseLogger) BuildLog(level Level, format string) string {
	format = fmt.Sprintf(loggerFormat, level.String(), format)
	return format
}

func (l *baseLogger) BuildErrorLog(level Level, err error, format string) string {
	format = fmt.Sprintf(errorLoggerFormat, level.String(), format, err)
	return format
}

func (l *baseLogger) init(wb writeBuilder, level Level) {
	l.builder = wb
	l.Writer = wb
	l.level = level
}

func (l *baseLogger) AppendRowTerminator(format string) string {
	if !strings.HasSuffix(format, defaultTerminator) {
		format = format + defaultTerminator
	}
	return format
}

type builder interface {
	BuildLog(level Level, format string) string
	BuildErrorLog(level Level, err error, format string) string
	AppendRowTerminator(format string) string
}

type writeBuilder interface {
	Writer
	builder
}

func checkLogLevel(level Level, currentLevel Level) bool {
	return currentLevel >= level
}
