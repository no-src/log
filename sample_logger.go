package log

import (
	"github.com/no-src/log/formatter"
	"github.com/no-src/log/sample"
)

type sampleLogger struct {
	sampleFunc sample.SampleFunc
	rate       float64
	logger     Logger
}

// NewDefaultSampleLogger get a sample logger with custom sample rate
func NewDefaultSampleLogger(logger Logger, sampleRate float64) Logger {
	return NewSampleLogger(logger, sample.DefaultSampleFunc, sampleRate)
}

// NewSampleLogger get a sample logger with custom sample rate and sample function
func NewSampleLogger(logger Logger, sampleFunc sample.SampleFunc, sampleRate float64) Logger {
	l := &sampleLogger{
		sampleFunc: sampleFunc,
		rate:       sampleRate,
		logger:     logger,
	}
	return l
}

func (l *sampleLogger) Debug(format string, args ...any) {
	if l.sample() {
		l.logger.Debug(format, args...)
	}
}

func (l *sampleLogger) Info(format string, args ...any) {
	if l.sample() {
		l.logger.Info(format, args...)
	}
}

func (l *sampleLogger) Warn(format string, args ...any) {
	if l.sample() {
		l.logger.Warn(format, args...)
	}
}

func (l *sampleLogger) Error(err error, format string, args ...any) {
	if l.sample() {
		l.logger.Error(err, format, args...)
	}
}

func (l *sampleLogger) ErrorIf(err error, format string, args ...any) error {
	return errorIf(l.Error, err, format, args...)
}

func (l *sampleLogger) Log(format string, args ...any) {
	l.logger.Log(format, args...)
}

func (l *sampleLogger) Close() error {
	return l.logger.Close()
}

func (l *sampleLogger) Write(p []byte) (n int, err error) {
	return l.logger.Write(p)
}

func (l *sampleLogger) sample() bool {
	return l.sampleFunc(l.rate)
}

func (l *sampleLogger) WithFormatter(f formatter.Formatter) Logger {
	l.logger.WithFormatter(f)
	return l
}

func (l *sampleLogger) WithTimeFormat(f string) Logger {
	l.logger.WithTimeFormat(f)
	return l
}
