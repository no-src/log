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

// Debug write the debug log
func Debug(format string, args ...interface{}) {
	defaultLogger.Debug(format, args...)
}

// Info write the info log
func Info(format string, args ...interface{}) {
	defaultLogger.Info(format, args...)
}

// Warn write the warn log
func Warn(format string, args ...interface{}) {
	defaultLogger.Warn(format, args...)
}

// Error write the error log
func Error(err error, format string, args ...interface{}) {
	defaultLogger.Error(err, format, args...)
}

// ErrorIf write the error log if err is not nil
func ErrorIf(err error, format string, args ...interface{}) {
	if err != nil {
		Error(err, format, args...)
	}
}

// Log write the log without level
func Log(format string, args ...interface{}) {
	defaultLogger.Log(format, args...)
}

// Close close the current logger
func Close() error {
	return defaultLogger.Close()
}

// DefaultLogger return the global default logger
func DefaultLogger() Logger {
	return defaultLogger
}

func init() {
	if defaultLogger == nil {
		defaultLogger = NewConsoleLogger(InfoLevel)
	}
}
