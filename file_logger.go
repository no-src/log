package log

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/no-src/log/formatter"
	"github.com/no-src/log/internal/cbool"
	"github.com/no-src/log/level"
	"github.com/no-src/log/option"
)

var (
	flushLogMsg = logMsg{flush: true}
	closeLogMsg = logMsg{closed: true}
	newWriter   = bufio.NewWriter
	stat        = os.Stat
	isNotExist  = os.IsNotExist
	mkdirAll    = os.MkdirAll
	create      = os.Create
	openFile    = os.OpenFile
	now         = time.Now
)

type fileLogger struct {
	baseLogger

	opt         option.FileLoggerOption
	in          chan logMsg
	writer      *bufio.Writer
	initialized bool
	closed      *cbool.CBool
	close       chan struct{} // the log is closed, and wait to write all the logs
	mu          sync.Mutex    // avoid data race for writer
	date        time.Time
	file        *os.File
}

type logMsg struct {
	log    []byte
	closed bool
	flush  bool
}

// NewFileLogger get a default file logger, auto flush logs to file per 3 seconds by default
func NewFileLogger(lvl level.Level, logDir string, filePrefix string) (Logger, error) {
	return NewFileLoggerWithAutoFlush(lvl, logDir, filePrefix, true, time.Second*3)
}

// NewFileLoggerWithAutoFlush get a file logger
func NewFileLoggerWithAutoFlush(lvl level.Level, logDir string, filePrefix string, autoFlush bool, flushInterval time.Duration) (Logger, error) {
	return NewFileLoggerWithOption(option.NewFileLoggerOption(lvl, logDir, filePrefix, autoFlush, flushInterval, false))
}

// NewFileLoggerWithOption get a file logger with option
func NewFileLoggerWithOption(opt option.FileLoggerOption) (Logger, error) {
	logger := &fileLogger{
		opt:    opt,
		in:     make(chan logMsg, 100),
		close:  make(chan struct{}, 1),
		mu:     sync.Mutex{},
		closed: cbool.New(false),
	}
	// init baseLogger
	logger.baseLogger.init(logger, opt.Level, true)
	// init fileLogger
	err := logger.init()
	return logger, err
}

func (l *fileLogger) Close() error {
	// stop a new log to write
	l.closed.Set(true)
	// send a closed message
	l.in <- closeLogMsg
	// wait to receive a close finished message
	<-l.close
	return nil
}

func (l *fileLogger) init() error {
	if err := l.initFile(true); err != nil {
		return err
	}
	l.start()
	l.startAutoFlush()
	return nil
}

func (l *fileLogger) initFile(splitDate bool) error {
	if !splitDate {
		return nil
	}
	date := now()
	timeFormat := "20060102"
	if l.date.Format(timeFormat) == date.Format(timeFormat) {
		return nil
	}
	// reset initialized
	l.initialized = false

	logDir := filepath.Clean(l.opt.LogDir)
	prefix := strings.TrimSpace(l.opt.FilePrefix)
	if len(prefix) > 0 {
		prefix = filepath.Clean(prefix)
	}
	logFile := logDir + "/" + prefix + date.Format(timeFormat) + ".log"

	_, err := stat(logDir)
	if isNotExist(err) {
		err = mkdirAll(logDir, 0766)
		if err != nil {
			l.innerLog("init file logger err, can't create the log dir. %s", err)
			return err
		}
	}
	_, err = stat(logFile)
	if isNotExist(err) {
		_, err = create(logFile)
		if err != nil {
			l.innerLog("init file logger err, can't create the log file. %s", err)
			return err
		}
	}
	f, err := openFile(logFile, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		l.innerLog("init file logger err, can't open the log file. %s", err)
		return err
	}

	if l.writer != nil {
		l.writer.Flush()
	}

	l.writer = newWriter(f)
	l.initialized = true
	l.date = date

	if l.file != nil {
		if err = l.file.Close(); err != nil {
			l.innerLog("close file error => %s %v", l.file.Name(), err)
		}
	}
	l.file = f
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
			l.file.Close()
		}
		l.close <- struct{}{}
	} else if msg.flush && l.initialized && l.writer != nil && l.writer.Buffered() > 0 {
		// received a flush message, flush logs to file
		if err := l.writer.Flush(); err != nil {
			l.innerLog("file logger flush log error. %s", err)
		}
	} else if l.initialized && l.writer != nil && len(msg.log) > 0 {
		if err := l.initFile(l.opt.SplitDate); err != nil {
			l.innerLog("init log file error. %s", err)
			return
		}
		if _, err := l.writer.Write(msg.log); err != nil {
			l.innerLog("file logger write log error. %s", err)
		}
	}
}

func (l *fileLogger) innerLog(format string, args ...any) {
	fmt.Printf(format+"\n", args...)
}

// startAutoFlush start to auto flush log to file per flushInterval
// if buffered is empty by 10 times checked, to delay wait next check
func (l *fileLogger) startAutoFlush() {
	if !l.opt.AutoFlush || l.opt.FlushInterval <= 0 {
		return
	}
	go func() {
		wait := l.opt.FlushInterval
		nop := 0
		delayCheckCount := 10
		for {
			<-time.After(wait)
			if l.closed.Get() {
				return
			}
			l.mu.Lock()
			buffered := l.writer.Buffered()
			l.mu.Unlock()
			if buffered > 0 {
				l.in <- flushLogMsg
				wait = l.opt.FlushInterval
				nop = 0
			} else {
				nop++
				if nop >= delayCheckCount {
					wait += l.opt.FlushInterval
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
	if !l.closed.Get() {
		l.in <- logMsg{
			log:    cp,
			closed: false,
		}
		return pLen, nil
	}
	return 0, errors.New("file logger is uninitialized or closed")
}

func (l *fileLogger) WithFormatter(f formatter.Formatter) Logger {
	l.setFormatter(f)
	return l
}

func (l *fileLogger) WithTimeFormat(f string) Logger {
	l.setTimeFormat(f)
	return l
}
