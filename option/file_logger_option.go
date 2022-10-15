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
