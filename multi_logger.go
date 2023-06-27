package log

import (
	"github.com/no-src/log/formatter"
)

type multiLogger struct {
	loggers []Logger
}

// NewMultiLogger get a multi logger, write log to multiple loggers
func NewMultiLogger(loggers ...Logger) Logger {
	logger := &multiLogger{}
	for _, l := range loggers {
		if l != nil {
			logger.loggers = append(logger.loggers, l)
		}
	}
	return logger
}

func (l *multiLogger) Debug(format string, args ...any) {
	for _, logger := range l.loggers {
		logger.Debug(format, args...)
	}
}

func (l *multiLogger) Info(format string, args ...any) {
	for _, logger := range l.loggers {
		logger.Info(format, args...)
	}
}

func (l *multiLogger) Warn(format string, args ...any) {
	for _, logger := range l.loggers {
		logger.Warn(format, args...)
	}
}

func (l *multiLogger) Error(err error, format string, args ...any) {
	for _, logger := range l.loggers {
		logger.Error(err, format, args...)
	}
}

func (l *multiLogger) Log(format string, args ...any) {
	for _, logger := range l.loggers {
		logger.Log(format, args...)
	}
}

func (l *multiLogger) Close() error {
	var err error
	for _, logger := range l.loggers {
		closeErr := logger.Close()
		if closeErr != nil {
			err = closeErr
		}
	}
	return err
}

func (l *multiLogger) Write(p []byte) (n int, err error) {
	n = len(p)
	for _, logger := range l.loggers {
		nn, writeErr := logger.Write(p)
		if writeErr != nil {
			err = writeErr
			n = nn
		}
	}
	return n, err
}

func (l *multiLogger) WithFormatter(f formatter.Formatter) Logger {
	for _, logger := range l.loggers {
		logger.WithFormatter(f)
	}
	return l
}

func (l *multiLogger) WithTimeFormat(f string) Logger {
	for _, logger := range l.loggers {
		logger.WithTimeFormat(f)
	}
	return l
}
