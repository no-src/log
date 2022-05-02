package log

import (
	"testing"
	"time"
)

func TestFileLogger(t *testing.T) {
	fileLogger, err := NewFileLogger(DebugLevel, "./logs", "ns")
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(fileLogger)
	defer Close()
	TestLogs(t)
}

func TestFileLogger_WithAutoFlush(t *testing.T) {
	wait := time.Second * 1
	autoFlushFileLogger, err := NewFileLoggerWithAutoFlush(DebugLevel, "./logs", "ns", true, wait)
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(autoFlushFileLogger)
	TestLogs(t)
	<-time.After(wait + time.Second)
}

func TestFileLogger_WithAutoFlushWithCloseWhenWrite(t *testing.T) {
	wait := time.Second * 1
	autoFlushFileLogger, err := NewFileLoggerWithAutoFlush(DebugLevel, "./logs", "ns", true, wait)
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(autoFlushFileLogger)
	go func() {
		Close()
	}()
	TestLogs(t)
	<-time.After(wait + time.Second)
}

func TestFileLogger_WithAutoFlushWithFlushDelay(t *testing.T) {
	wait := time.Millisecond * 10
	autoFlushFileLogger, err := NewFileLoggerWithAutoFlush(DebugLevel, "./logs", "ns", true, wait)
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(autoFlushFileLogger)
	<-time.After(wait * 20)
}

func TestConsoleLoggerAndFileLogger(t *testing.T) {
	fileLogger, err := NewFileLogger(DebugLevel, "./multi_logs", "ns")
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(NewMultiLogger(NewConsoleLogger(DebugLevel), fileLogger))
	defer Close()
	TestLogs(t)
}

func init() {
	initFileLoggerMock()
}
