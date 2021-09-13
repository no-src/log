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
	level  Level
	logDir string
	in     chan string
	writer *bufio.Writer
	// does file logger initialized
	initialized bool
	filePrefix  string
}

// NewFileLogger get a file logger
func NewFileLogger(level Level, logDir string, filePrefix string) Logger {
	logger := &fileLogger{
		level:      level,
		logDir:     logDir,
		in:         make(chan string, 10),
		filePrefix: filePrefix,
	}
	logger.baseLogger.Writer = logger
	logger.init()
	return logger
}

// Log write a format log to file
func (l *fileLogger) Log(format string, args ...interface{}) {
	if l.initialized {
		format = fmt.Sprintf("[%s] ", time.Now().Format("2006-01-02 15:04:05")) + format
		format = fmt.Sprintf(format, args...)
		format = appendRowTerminator(format)
		l.in <- format
	}
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
	log := <-l.in
	if l.initialized && l.writer != nil && len(log) > 0 {
		l.writer.WriteString(log)
		l.writer.Flush()
	}
}

func (l *fileLogger) innerLog(format string, args ...interface{}) {
	fmt.Printf(format+"\n", args...)
}
