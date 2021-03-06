# log

[![Chat](https://img.shields.io/discord/936876326722363472)](https://discord.gg/n47twC6Kcb)
[![Build](https://img.shields.io/github/workflow/status/no-src/log/Go)](https://github.com/no-src/log/actions)
[![License](https://img.shields.io/github/license/no-src/log)](https://github.com/no-src/log/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/no-src/log.svg)](https://pkg.go.dev/github.com/no-src/log)
[![Go Report Card](https://goreportcard.com/badge/github.com/no-src/log)](https://goreportcard.com/report/github.com/no-src/log)
[![codecov](https://codecov.io/gh/no-src/log/branch/main/graph/badge.svg?token=8Q20UR86EW)](https://codecov.io/gh/no-src/log)
[![Release](https://img.shields.io/github/v/release/no-src/log)](https://github.com/no-src/log/releases)

## Installation

```bash
go get -u github.com/no-src/log
```

## Quick Start

Current support following loggers

- empty logger
- console logger
- file logger
- multi logger
- sample logger

For example, init a file logger, to write logs.

```go
package main

import (
	"errors"

	"github.com/no-src/log"
	"github.com/no-src/log/level"
)

func main() {
	// init default logger
	if logger, err := log.NewFileLogger(level.DebugLevel, "./logs", ""); err == nil {
		log.InitDefaultLoggerWithSample(logger, 0.6)
	} else {
		log.Error(err, "init file logger error")
	}
	defer log.Close()

	// use default logger
	log.Debug("%s %s, test debug log", "hello", "world")
	log.Info("%s %s, test info log", "hello", "world")
	log.Warn("%s %s, test warn log", "hello", "world")
	log.Error(errors.New("log err"), "%s %s, test error log", "hello", "world")
	log.ErrorIf(errors.New("log err"), "%s %s, test error log", "hello", "world")
	log.Log("%s %s, test log log", "hello", "world")

	// use default logger by random sampling
	log.DebugSample("[sample] %s %s, test debug log", "hello", "world")
	log.InfoSample("[sample] %s %s, test info log", "hello", "world")
	log.WarnSample("[sample] %s %s, test warn log", "hello", "world")
	log.ErrorSample(errors.New("log err"), "[sample] %s %s,test error log", "hello", "world")
	log.ErrorIfSample(errors.New("log err from ErrorIfSample"), "[sample] %s %s, test error log", "hello", "world")
}
```