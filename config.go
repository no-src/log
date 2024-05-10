package log

import (
	"time"

	"github.com/no-src/log/level"
	"github.com/no-src/log/option"
)

type config struct {
	Loggers []loggerConfig `yaml:"loggers"`
}

type loggerConfig struct {
	// common fields
	Name       string  `yaml:"name"`
	Type       string  `yaml:"type"`
	Level      string  `yaml:"level"`
	Format     string  `yaml:"format"`
	TimeFormat string  `yaml:"time-format"`
	Sample     float64 `yaml:"sample"`

	// file logger fields
	LogDir            string        `yaml:"log-dir"`
	LogFilePrefix     string        `yaml:"log-file-prefix"`
	AutoFlush         bool          `yaml:"auto-flush"`
	AutoFlushInterval time.Duration `yaml:"auto-flush-interval"`
	SplitByDate       bool          `yaml:"split-by-date"`
}

func toFileLoggerOption(logConf loggerConfig) (opt option.FileLoggerOption, err error) {
	lvl, err := level.ParseLevel(logConf.Level)
	if err != nil {
		return opt, err
	}
	opt = option.NewFileLoggerOption(lvl, logConf.LogDir, logConf.LogFilePrefix, logConf.AutoFlush, logConf.AutoFlushInterval, logConf.SplitByDate)
	return opt, nil
}
