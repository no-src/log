package log

import (
	"sync"
	"testing"
	"time"

	"github.com/no-src/log/formatter"
	"github.com/no-src/log/level"
	"github.com/no-src/log/option"
)

func TestFileLogger(t *testing.T) {
	testCases := []struct {
		name        string
		formatter   string
		concurrency bool
		timeFormat  string
	}{
		{"TextFormatter", formatter.TextFormatter, false, testTimeFormat},
		{"JsonFormatter", formatter.JsonFormatter, false, testTimeFormat},
		{"TextFormatter Concurrency", formatter.TextFormatter, true, ""},
		{"JsonFormatter Concurrency", formatter.JsonFormatter, true, ""},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fLogger, err := NewFileLogger(level.DebugLevel, "./logs", "ns_"+tc.formatter)
			if err != nil {
				t.Fatal(err)
			}
			InitDefaultLogger(fLogger.WithFormatter(formatter.New(tc.formatter)).WithTimeFormat(tc.timeFormat))
			defer Close()
			if tc.concurrency {
				testLogsConcurrency(t, "TestFileLogger")
			} else {
				testLogs(t)
			}
		})
	}
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
	var wg sync.WaitGroup
	wg.Add(1)
	wait := time.Second * 1
	autoFlushFileLogger, err := NewFileLoggerWithAutoFlush(level.DebugLevel, "./logs", "ns", true, wait)
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(autoFlushFileLogger)
	go func() {
		Close()
		wg.Done()
	}()
	testLogs(t)
	// wait to close and stop the auto flush
	<-time.After(wait + time.Second)
	wg.Wait()
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
	fLogger, err := NewFileLogger(level.DebugLevel, "./multi_logs", "ns")
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(NewMultiLogger(NewConsoleLogger(level.DebugLevel), fLogger))
	defer Close()
	testLogs(t)
}

func TestFileLogger_WithSplitDate(t *testing.T) {
	fLogger, err := NewFileLoggerWithOption(option.FileLoggerOption{
		Level:      level.DebugLevel,
		LogDir:     "./split_date_logs",
		FilePrefix: "ns",
		SplitDate:  true,
	})
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(fLogger)
	defer Close()
	testLogs(t)
}

func TestFileLogger_InitFileError(t *testing.T) {
	_, err := NewFileLogger(level.DebugLevel, "/", "")
	if err == nil {
		t.Errorf("expect to get an error but get nil")
	}
}

func TestFileLogger_WithMultiInitFile(t *testing.T) {
	fLogger, err := NewFileLogger(level.DebugLevel, "./multi_init_file_logs", "ns")
	if err != nil {
		t.Fatal(err)
	}
	now = tomorrow
	defer func() {
		now = time.Now
	}()
	fLogger.(*fileLogger).initFile(true)
	InitDefaultLogger(fLogger)
	defer Close()
}

func tomorrow() time.Time {
	return time.Now().Add(time.Hour * 24)
}
