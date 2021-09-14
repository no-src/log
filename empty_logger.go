package log

type emptyLogger struct {
}

// NewEmptyLogger get an empty logger, there is nothing to do
func NewEmptyLogger() Logger {
	logger := &emptyLogger{}
	return logger
}

func (l *emptyLogger) Debug(format string, args ...interface{}) {

}

func (l *emptyLogger) Info(format string, args ...interface{}) {

}

func (l *emptyLogger) Warn(format string, args ...interface{}) {

}

func (l *emptyLogger) Error(err error, format string, args ...interface{}) {

}

func (l *emptyLogger) Log(format string, args ...interface{}) {

}

func (l *emptyLogger) Close() error {
	return nil
}
