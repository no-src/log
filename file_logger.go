package log

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"time"
)

type fileLogger struct {
	baseLogger
	logDir      string
	in          chan logMsg
	writer      *bufio.Writer
	initialized bool
	filePrefix  string
	closed      bool
	close       chan bool // the log is closed, and wait to write all the logs
}

type logMsg struct {
	log    string
	closed bool
}

// NewFileLogger get a file logger
func NewFileLogger(level Level, logDir string, filePrefix string) Logger {
	logger := &fileLogger{
		logDir:     logDir,
		in:         make(chan logMsg, 10),
		close:      make(chan bool, 1),
		filePrefix: filePrefix,
	}
	logger.baseLogger.Writer = logger
	logger.baseLogger.level = level
	logger.init()
	return logger
}

// Log write a format log to file
func (l *fileLogger) Log(format string, args ...interface{}) {
	if l.initialized && !l.closed {
		format = fmt.Sprintf("[%s] ", time.Now().Format("2006-01-02 15:04:05")) + format
		format = fmt.Sprintf(format, args...)
		format = appendRowTerminator(format)
		l.in <- logMsg{
			log:    format,
			closed: false,
		}
	}
}

func (l *fileLogger) Close() error {
	// stop a new log to write
	l.closed = true
	// send a closed message
	l.in <- logMsg{closed: true}
	// wait to receive a close finished message
	<-l.close
	return nil
}

func (l *fileLogger) init() error {
	logDir := filepath.Clean(l.logDir)
	logFile := logDir + "/" + filepath.Clean(l.filePrefix) + time.Now().Format("20060102") + ".log"

	_, err := os.Stat(logDir)
	if os.IsNotExist(err) {
		err = os.MkdirAll(logDir, 0666)
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
	// if received a closed message, flush logs to file, and notify close finished.
	if msg.closed {
		if l.initialized && l.writer != nil {
			l.writer.Flush()
		}
		l.close <- true
		return
	}
	if l.initialized && l.writer != nil && len(msg.log) > 0 {
		l.writer.WriteString(msg.log)
	}
}

func (l *fileLogger) innerLog(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
