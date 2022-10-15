//go:build !no_mock

package log

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"
	"testing"
	"time"

	"github.com/no-src/log/level"
	"github.com/no-src/log/option"
)

func TestFileLogger_WithNotExistFile(t *testing.T) {
	isNotExist = isNotExistAlwaysTrueMock
	defer func() {
		initOSMock()
	}()
	fileLogger, err := NewFileLogger(level.DebugLevel, "./logs", "ns")
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(fileLogger)
	defer Close()
	testLogs(t)
}

func TestFileLogger_WithCreateLogDirError(t *testing.T) {
	isNotExist = isNotExistAlwaysTrueMock
	mkdirAll = mkdirAllErrorMock
	defer func() {
		initOSMock()
	}()
	_, err := NewFileLogger(level.DebugLevel, "./logs", "ns")
	if err == nil {
		t.Fatal("create file logger expect to get an error but get nil")
	}
}

func TestFileLogger_WithCreateLogFileError(t *testing.T) {
	isNotExist = isNotExistAlwaysTrueMock
	create = createErrorMock
	defer func() {
		initOSMock()
	}()
	_, err := NewFileLogger(level.DebugLevel, "./logs", "ns")
	if err == nil {
		t.Fatal("create file logger expect to get an error but get nil")
	}
}

func TestFileLogger_WithOpenLogFileError(t *testing.T) {
	isNotExist = isNotExistAlwaysTrueMock
	openFile = openFileErrorMock
	defer func() {
		initOSMock()
	}()
	_, err := NewFileLogger(level.DebugLevel, "./logs", "ns")
	if err == nil {
		t.Fatal("create file logger expect to get an error but get nil")
	}
}

func TestFileLogger_WithAutoFlushWithWriteError(t *testing.T) {
	initNewWriterErrorMock()
	defer func() {
		initNewWriterMock()
	}()
	wait := time.Second * 1
	autoFlushFileLogger, err := NewFileLoggerWithAutoFlush(level.DebugLevel, "./logs", "ns", true, wait)
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(autoFlushFileLogger)
	testLogs(t)
	// more than the defaultBufSize 4096
	Debug(strings.Repeat("hello golang, hello gopher", 500))
	<-time.After(wait + time.Second)
}

func TestFileLogger_WithAutoFlushWithWriteErrorAndNoAutoFlush(t *testing.T) {
	initNewWriterErrorMock()
	defer func() {
		initNewWriterMock()
	}()
	wait := time.Second * 1
	autoFlushFileLogger, err := NewFileLoggerWithAutoFlush(level.DebugLevel, "./logs", "ns", false, wait)
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(autoFlushFileLogger)
	testLogs(t)
	<-time.After(wait + time.Second)
}

func TestFileLogger_WithMultiInitFile(t *testing.T) {
	defer func() {
		initTimeMock()
	}()
	fLogger, err := NewFileLogger(level.DebugLevel, "./multi_init_file_logs", "ns")
	if err != nil {
		t.Fatal(err)
	}
	now = nowMock
	fLogger.(*fileLogger).initFile(true)
	InitDefaultLogger(fLogger)
	defer Close()
}

func TestFileLogger_WithSplitDateError(t *testing.T) {
	defer func() {
		initTimeMock()
		initOSMock()
	}()
	fLogger, err := NewFileLoggerWithOption(option.FileLoggerOption{
		Level:      level.DebugLevel,
		LogDir:     "./split_date_logs_error",
		FilePrefix: "ns",
		SplitDate:  true,
	})
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(fLogger)
	defer Close()
	isNotExist = isNotExistAlwaysTrueMock
	mkdirAll = mkdirAllErrorMock
	now = nowMock
	testLogs(t)
}

func init() {
	initFileLoggerMock()
}

func initFileLoggerMock() {
	initNewWriterMock()
	initOSMock()
}

func initNewWriterMock() {
	newWriter = func(w io.Writer) *bufio.Writer {
		return bufio.NewWriter(NewEmptyLogger())
	}
}

func initNewWriterErrorMock() {
	newWriter = func(w io.Writer) *bufio.Writer {
		return bufio.NewWriter(errWriter{})
	}
}

type errWriter struct {
}

func (w errWriter) Write(p []byte) (n int, err error) {
	return 0, errors.New("the Write error test")
}

func initOSMock() {
	isNotExist = isNotExistAlwaysFalseMock
	mkdirAll = mkdirAllMock
	create = createMock
	openFile = openFileMock
}

func initTimeMock() {
	now = time.Now
}

func isNotExistAlwaysFalseMock(err error) bool {
	return false
}

func isNotExistAlwaysTrueMock(err error) bool {
	return true
}

func mkdirAllMock(path string, perm os.FileMode) error {
	return nil
}

func mkdirAllErrorMock(path string, perm os.FileMode) error {
	return errors.New("the MkdirAll error test")
}

func createMock(name string) (*os.File, error) {
	return nil, nil
}

func createErrorMock(name string) (*os.File, error) {
	return nil, errors.New("the Create error test")
}

func openFileMock(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.Stdout, nil
}

func openFileErrorMock(name string, flag int, perm os.FileMode) (*os.File, error) {
	return nil, errors.New("the OpenFile error test")
}

func nowMock() time.Time {
	return time.Now().Add(time.Hour * 24)
}
