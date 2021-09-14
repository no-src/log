package log

var defaultLogger Logger

// InitDefaultLogger init a default logger
// if InitDefaultLogger is not called, default is consoleLogger with InfoLevel
func InitDefaultLogger(logger Logger) {
	defaultLogger = logger
	if defaultLogger == nil {
		defaultLogger = NewEmptyLogger()
	}
}

func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

func Warn(format string, args ...interface{}) {
	defaultLogger.Warn(format, args...)
}

func Error(err error, format string, args ...interface{}) {
	defaultLogger.Error(err, format, args...)
}

func Log(format string, args ...interface{}) {
	defaultLogger.Log(format, args...)
}

func Close() error {
	return defaultLogger.Close()
}

func init() {
	if defaultLogger == nil {
		defaultLogger = NewConsoleLogger(InfoLevel)
	}
}
