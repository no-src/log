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
	InitDefaultLogger(NewConsoleLogger(DebugLevel))
	TestLogs(t)
}

func TestEmptyLogger(t *testing.T) {
	InitDefaultLogger(NewEmptyLogger())
	TestLogs(t)
}

func TestFileLogger(t *testing.T) {
	InitDefaultLogger(NewFileLogger(DebugLevel, "./logs", "ns"))
	TestLogs(t)
}

func TestMultiLogger(t *testing.T) {
	InitDefaultLogger(NewMultiLogger(NewConsoleLogger(DebugLevel), NewFileLogger(DebugLevel, "./multi_logs", "ns")))
	TestLogs(t)
}
