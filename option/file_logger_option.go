package option

import (
	"time"

	"github.com/no-src/log/level"
)

// FileLoggerOption the option of the fileLogger
type FileLoggerOption struct {
	Level         level.Level
	LogDir        string
	FilePrefix    string
	AutoFlush     bool
	FlushInterval time.Duration
	SplitDate     bool
}

// NewFileLoggerOption returns an instance of the FileLoggerOption
func NewFileLoggerOption(lvl level.Level, logDir string, filePrefix string, autoFlush bool, flushInterval time.Duration, splitDate bool) FileLoggerOption {
	return FileLoggerOption{
		Level:         lvl,
		LogDir:        logDir,
		FilePrefix:    filePrefix,
		AutoFlush:     autoFlush,
		FlushInterval: flushInterval,
		SplitDate:     splitDate,
	}
}
