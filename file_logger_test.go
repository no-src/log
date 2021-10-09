//go:build file_logger
// +build file_logger

package log

import (
	"testing"
	"time"
)

func TestFileLogger(t *testing.T) {
	fileLogger, err := NewFileLogger(DebugLevel, "./logs", "ns")
	if err != nil {
		t.FailNow()
	}
	InitDefaultLogger(fileLogger)
	defer Close()
	TestLogs(t)
}

func TestFileLoggerWithAutoFlush(t *testing.T) {
	wait := time.Second * 1
	autoFlushFileLogger, err := NewFileLoggerWithAutoFlush(DebugLevel, "./logs", "ns", true, wait)
	if err != nil {
		t.FailNow()
	}
	InitDefaultLogger(autoFlushFileLogger)
	TestLogs(t)
	<-time.After(wait + time.Second)
}

func TestConsoleLoggerAndFileLogger(t *testing.T) {
	fileLogger, err := NewFileLogger(DebugLevel, "./multi_logs", "ns")
	if err != nil {
		t.FailNow()
	}
	InitDefaultLogger(NewMultiLogger(NewConsoleLogger(DebugLevel), fileLogger))
	defer Close()
	TestLogs(t)
}
