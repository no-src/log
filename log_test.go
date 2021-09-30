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
	Log("%s,test log log again", "world")
}

func TestConsoleLogger(t *testing.T) {
	InitDefaultLogger(NewConsoleLogger(DebugLevel))
	defer Close()
	TestLogs(t)
}

func TestEmptyLogger(t *testing.T) {
	InitDefaultLogger(NewEmptyLogger())
	defer Close()
	TestLogs(t)
}

func TestMultiLogger(t *testing.T) {
	InitDefaultLogger(NewMultiLogger(NewConsoleLogger(DebugLevel)))
	defer Close()
	TestLogs(t)
}

func TestMinLogLevel(t *testing.T) {
	InitDefaultLogger(NewConsoleLogger(InfoLevel))
	defer Close()
	TestLogs(t)
}
