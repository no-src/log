package log

import (
	"fmt"
	"sync"

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
	timeFormat string
	optMu      sync.RWMutex // protect Option
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
	format = formatter.AppendRowTerminator(format)
	if len(args) > 0 {
		format = fmt.Sprintf(format, args...)
	}
	l.Write([]byte(format))
}

func (l *baseLogger) log(lvl level.Level, format string, args ...interface{}) {
	l.logWithErr(nil, lvl, format, args...)
}

func (l *baseLogger) logWithErr(err error, lvl level.Level, format string, args ...interface{}) {
	if checkLogLevel(l.lvl, lvl) {
		l.optMu.RLock()
		c := content.NewContent(lvl, err, l.appendTime, l.timeFormat, format, args...)
		f := l.f
		l.optMu.RUnlock()
		data, _ := f.Serialize(c)
		l.Log(string(data))
	}
}

// Close the default implementation of Writer.Close.
// Nothing is going to be done here, provide this default implementation to avoid infinite loop call and stack overflow if the real struct does not implement the Writer.Close.
// As mentioned above, a panic will happen => runtime: goroutine stack exceeds 1000000000-byte limit.
// So, keep it here.
func (l *baseLogger) Close() error {
	return nil
}

func (l *baseLogger) init(w Writer, lvl level.Level, appendTime bool) {
	l.Writer = w
	l.lvl = lvl
	l.f = formatter.Default()
	l.appendTime = appendTime
	l.setTimeFormat(content.DefaultLogTimeFormat())
}

func (l *baseLogger) setFormatter(f formatter.Formatter) {
	if f != nil {
		l.optMu.Lock()
		l.f = f
		l.optMu.Unlock()
	}
}

func (l *baseLogger) setTimeFormat(f string) {
	if len(f) == 0 {
		f = content.DefaultLogTimeFormat()
	}
	l.optMu.Lock()
	l.timeFormat = f
	l.optMu.Unlock()
}

func checkLogLevel(lvl level.Level, currentLevel level.Level) bool {
	return currentLevel >= lvl
}
