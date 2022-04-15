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
)

func TestFileLoggerWithNotExistFile(t *testing.T) {
	isNotExist = isNotExistAlwaysTrueMock
	defer func() {
		initOSMock()
	}()
	fileLogger, err := NewFileLogger(DebugLevel, "./logs", "ns")
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(fileLogger)
	defer Close()
	TestLogs(t)
}

func TestFileLoggerWithCreateLogDirError(t *testing.T) {
	isNotExist = isNotExistAlwaysTrueMock
	mkdirAll = mkdirAllErrorMock
	defer func() {
		initOSMock()
	}()
	_, err := NewFileLogger(DebugLevel, "./logs", "ns")
	if err == nil {
		t.Fatal("create file logger expect to get an error but get nil")
	}
}

func TestFileLoggerWithCreateLogFileError(t *testing.T) {
	isNotExist = isNotExistAlwaysTrueMock
	create = createErrorMock
	defer func() {
		initOSMock()
	}()
	_, err := NewFileLogger(DebugLevel, "./logs", "ns")
	if err == nil {
		t.Fatal("create file logger expect to get an error but get nil")
	}
}

func TestFileLoggerWithOpenLogFileError(t *testing.T) {
	isNotExist = isNotExistAlwaysTrueMock
	openFile = openFileErrorMock
	defer func() {
		initOSMock()
	}()
	_, err := NewFileLogger(DebugLevel, "./logs", "ns")
	if err == nil {
		t.Fatal("create file logger expect to get an error but get nil")
	}
}

func TestFileLoggerWithAutoFlushWithWriteError(t *testing.T) {
	initNewWriterErrorMock()
	defer func() {
		initNewWriterMock()
	}()
	wait := time.Second * 1
	autoFlushFileLogger, err := NewFileLoggerWithAutoFlush(DebugLevel, "./logs", "ns", true, wait)
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(autoFlushFileLogger)
	TestLogs(t)
	// more than the defaultBufSize 4096
	Debug(strings.Repeat("hello golang, hello gopher", 500))
	<-time.After(wait + time.Second)
}

func TestFileLoggerWithAutoFlushWithWriteErrorAndNoAutoFlush(t *testing.T) {
	initNewWriterErrorMock()
	defer func() {
		initNewWriterMock()
	}()
	wait := time.Second * 1
	autoFlushFileLogger, err := NewFileLoggerWithAutoFlush(DebugLevel, "./logs", "ns", false, wait)
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(autoFlushFileLogger)
	TestLogs(t)
	<-time.After(wait + time.Second)
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
	return nil, nil
}

func openFileErrorMock(name string, flag int, perm os.FileMode) (*os.File, error) {
	return nil, errors.New("the OpenFile error test")
}
