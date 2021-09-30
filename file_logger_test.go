//go:build file_logger
// +build file_logger

package log

import (
	"testing"
	"time"
)

func TestFileLogger(t *testing.T) {
	InitDefaultLogger(NewFileLogger(DebugLevel, "./logs", "ns"))
	defer Close()
	TestLogs(t)
}

func TestFileLoggerWithAutoFlush(t *testing.T) {
	wait := time.Second * 1
	InitDefaultLogger(NewFileLoggerWithAutoFlush(DebugLevel, "./logs", "ns", true, wait))
	TestLogs(t)
	<-time.After(wait + time.Second)
}

func TestConsoleLoggerAndFileLogger(t *testing.T) {
	InitDefaultLogger(NewMultiLogger(NewConsoleLogger(DebugLevel), NewFileLogger(DebugLevel, "./multi_logs", "ns")))
	defer Close()
	TestLogs(t)
}
