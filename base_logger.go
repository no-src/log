package log

// baseLogger Implement basic logger operation
type baseLogger struct {
	level Level // min log level
	Writer
}

func (l *baseLogger) Debug(format string, args ...interface{}) {
	if checkLogLevel(l.level, DebugLevel) {
		format = buildLog(DebugLevel, format)
		l.Writer.Log(format, args...)
	}
}

func (l *baseLogger) Info(format string, args ...interface{}) {
	if checkLogLevel(l.level, InfoLevel) {
		format = buildLog(InfoLevel, format)
		l.Writer.Log(format, args...)
	}
}

func (l *baseLogger) Warn(format string, args ...interface{}) {
	if checkLogLevel(l.level, WarnLevel) {
		format = buildLog(WarnLevel, format)
		l.Writer.Log(format, args...)
	}
}

func (l *baseLogger) Error(err error, format string, args ...interface{}) {
	if checkLogLevel(l.level, ErrorLevel) {
		format = buildErrorLog(ErrorLevel, err, format)
		l.Writer.Log(format, args...)
	}
}
