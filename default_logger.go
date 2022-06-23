package log

import (
	"github.com/no-src/log/level"
)

var (
	defaultLogger       Logger
	defaultSampleLogger Logger
)

const defaultSampleRate = 1

// InitDefaultLogger init a default logger
// if not specified, default is consoleLogger with InfoLevel, and default sample rate is 1
func InitDefaultLogger(logger Logger) {
	InitDefaultLoggerWithSample(logger, defaultSampleRate)
}

// InitDefaultLoggerWithSample init a default logger and sample logger
// if not specified, default is consoleLogger with InfoLevel, and default sample rate is 1
func InitDefaultLoggerWithSample(logger Logger, sampleRate float64) {
	defaultLogger = logger
	if defaultLogger == nil {
		defaultLogger = NewEmptyLogger()
	}
	defaultSampleLogger = NewDefaultSampleLogger(defaultLogger, sampleRate)
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
func ErrorIf(err error, format string, args ...interface{}) error {
	if err != nil {
		Error(err, format, args...)
	}
	return err
}

// DebugSample write the debug log by random sampling
func DebugSample(format string, args ...interface{}) {
	defaultSampleLogger.Debug(format, args...)
}

// InfoSample write the info log by random sampling
func InfoSample(format string, args ...interface{}) {
	defaultSampleLogger.Info(format, args...)
}

// WarnSample write the warn log by random sampling
func WarnSample(format string, args ...interface{}) {
	defaultSampleLogger.Warn(format, args...)
}

// ErrorSample write the error log by random sampling
func ErrorSample(err error, format string, args ...interface{}) {
	defaultSampleLogger.Error(err, format, args...)
}

// ErrorIfSample write the error log by random sampling if err is not nil
func ErrorIfSample(err error, format string, args ...interface{}) error {
	if err != nil {
		ErrorSample(err, format, args...)
	}
	return err
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
	InitDefaultLogger(NewConsoleLogger(level.InfoLevel))
}
