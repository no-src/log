package log

import (
	"testing"
	"time"

	"github.com/no-src/log/level"
)

func TestFileLogger(t *testing.T) {
	fileLogger, err := NewFileLogger(level.DebugLevel, "./logs", "ns")
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(fileLogger)
	defer Close()
	testLogs(t)
}

func TestFileLogger_WithAutoFlush(t *testing.T) {
	wait := time.Second * 1
	autoFlushFileLogger, err := NewFileLoggerWithAutoFlush(level.DebugLevel, "./logs", "ns", true, wait)
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(autoFlushFileLogger)
	testLogs(t)
	<-time.After(wait + time.Second)
}

func TestFileLogger_WithAutoFlushWithCloseWhenWrite(t *testing.T) {
	wait := time.Second * 1
	autoFlushFileLogger, err := NewFileLoggerWithAutoFlush(level.DebugLevel, "./logs", "ns", true, wait)
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(autoFlushFileLogger)
	go func() {
		Close()
	}()
	testLogs(t)
	<-time.After(wait + time.Second)
}

func TestFileLogger_WithAutoFlushWithFlushDelay(t *testing.T) {
	wait := time.Millisecond * 10
	autoFlushFileLogger, err := NewFileLoggerWithAutoFlush(level.DebugLevel, "./logs", "ns", true, wait)
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(autoFlushFileLogger)
	<-time.After(wait * 20)
}

func TestConsoleLoggerAndFileLogger(t *testing.T) {
	fileLogger, err := NewFileLogger(level.DebugLevel, "./multi_logs", "ns")
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(NewMultiLogger(NewConsoleLogger(level.DebugLevel), fileLogger))
	defer Close()
	testLogs(t)
}

func init() {
	initFileLoggerMock()
}
