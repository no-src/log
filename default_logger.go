package log

import (
	"sync"

	"github.com/no-src/log/level"
)

var (
	defaultLogger       Logger
	defaultSampleLogger Logger
	mu                  sync.RWMutex
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
	mu.Lock()
	defer mu.Unlock()
	defaultLogger = logger
	if defaultLogger == nil {
		defaultLogger = NewEmptyLogger()
	}
	defaultSampleLogger = NewDefaultSampleLogger(defaultLogger, sampleRate)
}

// Debug write the debug log
func Debug(format string, args ...any) {
	DefaultLogger().Debug(format, args...)
}

// Info write the info log
func Info(format string, args ...any) {
	DefaultLogger().Info(format, args...)
}

// Warn write the warn log
func Warn(format string, args ...any) {
	DefaultLogger().Warn(format, args...)
}

// Error write the error log
func Error(err error, format string, args ...any) {
	DefaultLogger().Error(err, format, args...)
}

// ErrorIf write the error log if err is not nil
func ErrorIf(err error, format string, args ...any) error {
	return DefaultLogger().ErrorIf(err, format, args...)
}

// DebugSample write the debug log by random sampling
func DebugSample(format string, args ...any) {
	DefaultSampleLogger().Debug(format, args...)
}

// InfoSample write the info log by random sampling
func InfoSample(format string, args ...any) {
	DefaultSampleLogger().Info(format, args...)
}

// WarnSample write the warn log by random sampling
func WarnSample(format string, args ...any) {
	DefaultSampleLogger().Warn(format, args...)
}

// ErrorSample write the error log by random sampling
func ErrorSample(err error, format string, args ...any) {
	DefaultSampleLogger().Error(err, format, args...)
}

// ErrorIfSample write the error log by random sampling if err is not nil
func ErrorIfSample(err error, format string, args ...any) error {
	return DefaultSampleLogger().ErrorIf(err, format, args...)
}

// Log write the log without level
func Log(format string, args ...any) {
	DefaultLogger().Log(format, args...)
}

// Close close the current logger
func Close() error {
	return DefaultLogger().Close()
}

// DefaultLogger return the global default logger
func DefaultLogger() Logger {
	mu.RLock()
	defer mu.RUnlock()
	return defaultLogger
}

// DefaultSampleLogger return the global default sample logger
func DefaultSampleLogger() Logger {
	mu.RLock()
	defer mu.RUnlock()
	return defaultSampleLogger
}

func init() {
	InitDefaultLogger(NewConsoleLogger(level.InfoLevel))
}
