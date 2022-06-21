package log

import (
	"fmt"
	"github.com/no-src/log/content"
	"github.com/no-src/log/formatter"
	"github.com/no-src/log/level"
)

// baseLogger Implement basic logger operation
type baseLogger struct {
	Writer

	lvl        level.Level // min log level
	f          formatter.Formatter
	appendTime bool
}

func (l *baseLogger) Debug(format string, args ...interface{}) {
	l.log(level.DebugLevel, format, args...)
}

func (l *baseLogger) Info(format string, args ...interface{}) {
	l.log(level.InfoLevel, format, args...)
}

func (l *baseLogger) Warn(format string, args ...interface{}) {
	l.log(level.WarnLevel, format, args...)
}

func (l *baseLogger) Error(err error, format string, args ...interface{}) {
	l.logWithErr(err, level.ErrorLevel, format, args...)
}

// Log write a format log
func (l *baseLogger) Log(format string, args ...interface{}) {
	format = fmt.Sprintf(formatter.AppendRowTerminator(format), args...)
	l.Write([]byte(format))
}

func (l *baseLogger) log(lvl level.Level, format string, args ...interface{}) {
	l.logWithErr(nil, lvl, format, args...)
}

func (l *baseLogger) logWithErr(err error, lvl level.Level, format string, args ...interface{}) {
	if checkLogLevel(l.lvl, lvl) {
		data, _ := l.f.Serialize(content.NewContent(lvl, format, args, err, l.appendTime))
		l.Log(string(data))
	}
}

func (l *baseLogger) Close() error {
	return nil
}

func (l *baseLogger) init(w Writer, lvl level.Level, appendTime bool) {
	l.Writer = w
	l.lvl = lvl
	l.f = formatter.Default()
	l.appendTime = appendTime
}

func checkLogLevel(lvl level.Level, currentLevel level.Level) bool {
	return currentLevel >= lvl
}
