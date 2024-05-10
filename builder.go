package log

import (
	"fmt"
	"os"
	"strings"

	"github.com/no-src/log/level"
	"gopkg.in/yaml.v3"
)

const (
	consoleLoggerType = "console"
	fileLoggerType    = "file"
	emptyLoggerType   = "empty"
)

// CreateLoggerFromConfig create a logger from config file
func CreateLoggerFromConfig(configFile string) (Logger, error) {
	data, err := os.ReadFile(configFile)
	if err != nil {
		return nil, err
	}
	var conf config
	err = yaml.Unmarshal(data, &conf)
	if err != nil {
		return nil, err
	}
	return createLoggers(conf.Loggers)
}

func createLoggers(configs []loggerConfig) (Logger, error) {
	var loggers []Logger
	for _, logConf := range configs {
		logger, err := createLogger(logConf)
		if err != nil {
			return nil, err
		}
		loggers = append(loggers, logger)
	}
	length := len(loggers)
	if length == 0 {
		return NewEmptyLogger(), nil
	} else if length == 1 {
		return loggers[0], nil
	} else {
		return NewMultiLogger(loggers...), nil
	}
}

func createLogger(logConf loggerConfig) (Logger, error) {
	var logger Logger
	loggerType := strings.ToLower(logConf.Type)
	switch loggerType {
	case consoleLoggerType:
		lvl, err := level.ParseLevel(logConf.Level)
		if err != nil {
			return nil, err
		}
		logger = NewConsoleLogger(lvl)
	case fileLoggerType:
		opt, err := toFileLoggerOption(logConf)
		if err != nil {
			return nil, err
		}
		fl, err := NewFileLoggerWithOption(opt)
		if err != nil {
			return nil, err
		}
		logger = fl
	case emptyLoggerType:
		logger = NewEmptyLogger()
	default:
		return nil, fmt.Errorf("unsupported logger type: %s", logConf.Type)
	}

	if loggerType != emptyLoggerType && logger != nil && logConf.Sample < 1 {
		logger = NewDefaultSampleLogger(logger, logConf.Sample)
	}
	return logger, nil
}
