package log

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type fileLogger struct {
	baseLogger
	logDir        string
	in            chan logMsg
	writer        *bufio.Writer
	initialized   bool
	filePrefix    string
	closed        bool
	close         chan bool // the log is closed, and wait to write all the logs
	autoFlush     bool
	flushInterval time.Duration
	mu            sync.Mutex // avoid data race for writer
}

type logMsg struct {
	log    []byte
	closed bool
	flush  bool
}

var (
	flushLogMsg = logMsg{flush: true}
	closeLogMsg = logMsg{closed: true}
)

// NewFileLogger get a file logger
func NewFileLogger(level Level, logDir string, filePrefix string) (Logger, error) {
	return NewFileLoggerWithAutoFlush(level, logDir, filePrefix, false, time.Duration(0))
}

// NewFileLoggerWithAutoFlush get a file logger
func NewFileLoggerWithAutoFlush(level Level, logDir string, filePrefix string, autoFlush bool, flushInterval time.Duration) (Logger, error) {
	logger := &fileLogger{
		logDir:        logDir,
		in:            make(chan logMsg, 100),
		close:         make(chan bool, 1),
		filePrefix:    filePrefix,
		autoFlush:     autoFlush,
		flushInterval: flushInterval,
		mu:            sync.Mutex{},
	}
	// init baseLogger
	logger.baseLogger.init(logger, level)
	// init fileLogger
	err := logger.init()
	return logger, err
}

// Log write a format log to file
func (l *fileLogger) Log(format string, args ...interface{}) {
	format = fmt.Sprintf("[%s] ", time.Now().Format("2006-01-02 15:04:05")) + format
	format = fmt.Sprintf(format, args...)
	format = l.builder.AppendRowTerminator(format)
	l.Write([]byte(format))
}

func (l *fileLogger) Close() error {
	// stop a new log to write
	l.closed = true
	// send a closed message
	l.in <- closeLogMsg
	// wait to receive a close finished message
	<-l.close
	return nil
}

func (l *fileLogger) init() error {
	logDir := filepath.Clean(l.logDir)
	logFile := logDir + "/" + filepath.Clean(l.filePrefix) + time.Now().Format("20060102") + ".log"

	_, err := os.Stat(logDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0766)
		if err != nil {
			l.innerLog("init file logger err, can't create the log dir. %s", err)
			return err
		}
	}
	_, err = os.Stat(logFile)
	if os.IsNotExist(err) {
		_, err = os.Create(logFile)
		if err != nil {
			l.innerLog("init file logger err, can't create the log file. %s", err)
			return err
		}
	}
	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		l.innerLog("init file logger err, can't open the log file. %s", err)
		return err
	}
	l.writer = bufio.NewWriter(f)
	l.initialized = true
	l.start()
	l.startAutoFlush()
	return nil
}

// start create a goroutine to receive channel messages and write log to file
func (l *fileLogger) start() {
	go func() {
		for {
			l.write()
		}
	}()
}

func (l *fileLogger) write() {
	msg := <-l.in
	l.mu.Lock()
	defer l.mu.Unlock()
	// if received a closed message, flush logs to file, and notify close finished.
	if msg.closed {
		if l.initialized && l.writer != nil {
			l.writer.Flush()
		}
		l.close <- true
	} else if msg.flush && l.initialized && l.writer != nil && l.writer.Buffered() > 0 {
		// received a flush message, flush logs to file
		if err := l.writer.Flush(); err != nil {
			l.innerLog("file logger flush log error. %s", err)
		}
	} else if l.initialized && l.writer != nil && len(msg.log) > 0 {
		if _, err := l.writer.Write(msg.log); err != nil {
			l.innerLog("file logger write log error. %s", err)
		}
	}
}

func (l *fileLogger) innerLog(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}

// startAutoFlush start to auto flush log to file per flushInterval
// if buffered is empty by 10 times checked, to delay wait next check
func (l *fileLogger) startAutoFlush() {
	if !l.autoFlush || l.flushInterval <= 0 {
		return
	}
	go func() {
		wait := l.flushInterval
		nop := 0
		delayCheckCount := 10
		for {
			<-time.After(wait)
			if l.closed || l.writer == nil {
				l.innerLog("auto flush stopped")
				return
			}
			l.mu.Lock()
			buffered := l.writer.Buffered()
			l.mu.Unlock()
			if buffered > 0 {
				l.in <- flushLogMsg
				wait = l.flushInterval
				nop = 0
			} else {
				nop++
				if nop >= delayCheckCount {
					wait += l.flushInterval
					nop = 0
				}
			}
		}
	}()
}

// Write see io.Writer
func (l *fileLogger) Write(p []byte) (n int, err error) {
	// copy data to avoid data race from caller
	pLen := len(p)
	if pLen == 0 {
		return 0, nil
	}
	cp := make([]byte, pLen)
	copy(cp, p)
	if l.initialized && !l.closed {
		l.in <- logMsg{
			log:    cp,
			closed: false,
		}
		return pLen, nil
	} else {
		return 0, errors.New("file logger is uninitialized or closed")
	}
}
