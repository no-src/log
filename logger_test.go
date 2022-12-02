package log

import (
	"errors"
	"testing"
	"time"

	"github.com/no-src/log/formatter"
	"github.com/no-src/log/internal/sync"
	"github.com/no-src/log/level"
)

var (
	concurrencyCount   = 3
	concurrencyTimeout = time.Second * 5
)

func testLogs(t *testing.T) {
	Debug("%s %s, test debug log", "hello", "world")
	Info("%s %s, test info log", "hello", "world")
	Warn("%s %s, test warn log", "hello", "world")
	Error(errors.New("log err"), "%s %s,test error log", "hello", "world")
	ErrorIf(errors.New("log err from ErrorIf"), "%s %s, test error log", "hello", "world")
	ErrorIf(nil, "%s %s, this error log will not be printed", "hello", "world")
	testSampleLogs()
	Log("%s %s, test log log", "hello", "world")
	Log("%s %s, test log log again", "hello", "world")
	DefaultLogger().Write([]byte(""))
	DefaultLogger().Write([]byte("hello logger"))
}

func testLogsConcurrency(t *testing.T, testName string) {
	wg := sync.WaitGroup{}
	wg.Add(concurrencyCount)
	for i := 0; i < concurrencyCount; i++ {
		go func() {
			testLogs(t)
			wg.Done()
		}()
	}
	if wg.WaitWithTimeout(concurrencyTimeout) {
		t.Errorf("[concurrency] %s timeout for %s", testName, concurrencyTimeout.String())
	}
}

func TestDefaultLogger(t *testing.T) {
	defer Close()
	testLogs(t)
}

func TestDefaultLogger_Concurrency(t *testing.T) {
	defer Close()
	testLogsConcurrency(t, "TestDefaultLogger_Concurrency")
}

func TestConsoleLogger(t *testing.T) {
	testCases := []struct {
		name        string
		formatter   string
		concurrency bool
	}{
		{"TextFormatter", formatter.TextFormatter, false},
		{"JsonFormatter", formatter.JsonFormatter, false},
		{"TextFormatter Concurrency", formatter.TextFormatter, true},
		{"JsonFormatter Concurrency", formatter.JsonFormatter, true},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InitDefaultLogger(NewConsoleLogger(level.DebugLevel).WithFormatter(formatter.New(tc.formatter)))
			defer Close()
			if tc.concurrency {
				testLogsConcurrency(t, "TestConsoleLogger")
			} else {
				testLogs(t)
			}
		})
	}
}

func TestConsoleLoggerWithBuffer(t *testing.T) {
	testCases := []struct {
		name      string
		formatter string
	}{
		{"TextFormatter", formatter.TextFormatter},
		{"JsonFormatter", formatter.JsonFormatter},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InitDefaultLogger(newConsoleLoggerWithBuffer(level.DebugLevel, true).WithFormatter(formatter.New(tc.formatter)))
			defer Close()
			testLogs(t)
		})
	}
}

func TestEmptyLogger(t *testing.T) {
	testCases := []struct {
		name      string
		formatter string
	}{
		{"TextFormatter", formatter.TextFormatter},
		{"JsonFormatter", formatter.JsonFormatter},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InitDefaultLogger(NewEmptyLogger().WithFormatter(formatter.New(tc.formatter)))
			defer Close()
			testLogs(t)
		})
	}
}

func TestMinLogLevel(t *testing.T) {
	InitDefaultLogger(NewConsoleLogger(level.InfoLevel))
	defer Close()
	testLogs(t)
}

func TestNilLogger(t *testing.T) {
	InitDefaultLogger(nil)
	defer Close()
	testLogs(t)
}

func TestBaseLogger_Close(t *testing.T) {
	InitDefaultLogger(newMinLogger())
	// call baseLogger.Close
	defer Close()
	testLogs(t)
}

func TestReadWriteLoggerConcurrency(t *testing.T) {
	go func() {
		for i := 0; i < 10; i++ {
			InitDefaultLogger(NewConsoleLogger(level.DebugLevel))
		}
	}()
	testLogsConcurrency(t, "TestReadWriteLoggerConcurrency")
	defer Close()
}

type minLogger struct {
	baseLogger
}

func newMinLogger() Logger {
	logger := &minLogger{}
	logger.init(logger, level.DebugLevel, true)
	return logger
}

func (l *minLogger) Write(p []byte) (n int, err error) {
	return len(p), nil
}

func (l *minLogger) WithFormatter(f formatter.Formatter) Logger {
	return l
}
