package log

import (
	"errors"
	"testing"

	"github.com/no-src/log/formatter"
	"github.com/no-src/log/level"
)

func TestLogs(t *testing.T) {
	Debug("%s, test debug log", "hello")
	Info("%s, test info log", "hello")
	Warn("%s, test warn log", "hello")
	Error(errors.New("log err"), "%s,test error log", "hello")
	ErrorIf(errors.New("log err from ErrorIf"), "%s, test error log", "hello")
	ErrorIf(nil, "%s, this error log will not be printed", "hello")
	testSampleLogs()
	Log("%s, test log log", "hello")
	Log("%s, test log log again", "world")
	DefaultLogger().Write([]byte(""))
	DefaultLogger().Write([]byte("hello logger"))
}

func TestConsoleLogger(t *testing.T) {
	InitDefaultLogger(NewConsoleLogger(level.DebugLevel))
	defer Close()
	TestLogs(t)
}

func TestEmptyLogger(t *testing.T) {
	InitDefaultLogger(NewEmptyLogger())
	defer Close()
	TestLogs(t)
}

func TestMinLogLevel(t *testing.T) {
	InitDefaultLogger(NewConsoleLogger(level.InfoLevel))
	defer Close()
	TestLogs(t)
}

func TestNilLogger(t *testing.T) {
	InitDefaultLogger(nil)
	defer Close()
	TestLogs(t)
}

func TestInitDefaultFormatter_TextFormatter(t *testing.T) {
	InitDefaultFormatter(formatter.TextFormatter)
	defer Close()
	TestLogs(t)
}

func TestInitDefaultFormatter_JsonFormatter(t *testing.T) {
	InitDefaultFormatter(formatter.JsonFormatter)
	defer Close()
	TestLogs(t)
}
