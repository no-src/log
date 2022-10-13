# log

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

- [Empty Logger](#empty-logger)
- [Console Logger](#console-logger)
- [File Logger](#file-logger)
- [Multi Logger](#multi-logger)
- [Sample Logger](#sample-logger)

For example, init a console logger, to write logs.

```go
package main

import (
	"errors"

	"github.com/no-src/log"
	"github.com/no-src/log/level"
)

func main() {
	// init console logger as default logger
	// replace the line of code with any logger you need
	log.InitDefaultLogger(log.NewConsoleLogger(level.DebugLevel))

	defer log.Close()

	text := "hello world"
	// use default logger to write logs
	log.Debug("%s, test debug log", text)
	log.Info("%s, test info log", text)
	log.Warn("%s, test warn log", text)
	log.Error(errors.New("log err"), "%s, test error log", text)
	log.ErrorIf(errors.New("log err"), "%s, test error log", text)
	log.Log("%s, test log log", text)
}
```

## Logger

### Empty Logger

Init empty logger as default logger.

```go
log.InitDefaultLogger(log.NewEmptyLogger())
```

### Console Logger

Init console logger as default logger.

```go
log.InitDefaultLogger(log.NewConsoleLogger(level.DebugLevel))
```

### File Logger

Init file logger as default logger.

```go
if logger, err := log.NewFileLogger(level.DebugLevel, "./logs", "default_"); err == nil {
    log.InitDefaultLogger(logger)
} else {
    log.Error(err, "init file logger error")
}
```

### Multi Logger

Init multi logger as default logger.

```go
if logger, err := log.NewFileLogger(level.DebugLevel, "./logs", "multi_"); err == nil {
    log.InitDefaultLogger(log.NewMultiLogger(log.NewConsoleLogger(level.DebugLevel), logger))
} else {
    log.Error(err, "init file logger error")
}
```

### Sample Logger

Init console logger as default logger and set the sample rate, default is `1`.

```go
log.InitDefaultLoggerWithSample(log.NewConsoleLogger(level.DebugLevel), 0.6)
```

Use default logger to write logs by random sampling.

```go
text := "hello world"
log.DebugSample("[sample] %s, test debug log", text)
log.InfoSample("[sample] %s, test info log", text)
log.WarnSample("[sample] %s, test warn log", text)
log.ErrorSample(errors.New("log err"), "[sample] %s, test error log", text)
log.ErrorIfSample(errors.New("log err from ErrorIfSample"), "[sample] %s, test error log", text)
```