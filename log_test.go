package log

import (
	"errors"
	"testing"
)

func TestLogs(t *testing.T) {
	Debug("%s,test debug log", "hello")
	Info("%s,test info log", "hello")
	Warn("%s,test warn log", "hello")
	Error(errors.New("log err"), "%s,test error log", "hello")
	Log("%s,test log log", "hello")
}

func TestConsoleLogger(t *testing.T) {
	consoleLogger := NewConsoleLogger(DebugLevel)
	defer consoleLogger.Close()
	InitDefaultLogger(consoleLogger)
	TestLogs(t)
}

func TestEmptyLogger(t *testing.T) {
	emptyLogger := NewEmptyLogger()
	defer emptyLogger.Close()
	InitDefaultLogger(emptyLogger)
	TestLogs(t)
}

func TestFileLogger(t *testing.T) {
	fileLogger := NewFileLogger(DebugLevel, "./logs", "ns")
	defer fileLogger.Close()
	InitDefaultLogger(fileLogger)
	TestLogs(t)
}

func TestMultiLogger(t *testing.T) {
	multiLogger := NewMultiLogger(NewConsoleLogger(DebugLevel), NewFileLogger(DebugLevel, "./multi_logs", "ns"))
	defer multiLogger.Close()
	InitDefaultLogger(multiLogger)
	TestLogs(t)
}

func TestMinLogLevel(t *testing.T) {
	multiLogger := NewMultiLogger(NewConsoleLogger(InfoLevel), NewFileLogger(InfoLevel, "./multi_logs", "lvl"))
	defer multiLogger.Close()
	InitDefaultLogger(multiLogger)
	TestLogs(t)
}
