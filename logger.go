package log

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
