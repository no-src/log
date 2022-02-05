# log

[![Chat](https://img.shields.io/discord/936876326722363472)](https://discord.gg/n47twC6Kcb)
[![Build](https://img.shields.io/github/workflow/status/no-src/log/Go)](https://github.com/no-src/log/actions)
[![License](https://img.shields.io/github/license/no-src/log)](https://github.com/no-src/log/blob/main/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/no-src/log.svg)](https://pkg.go.dev/github.com/no-src/log)
[![Release](https://img.shields.io/github/v/release/no-src/log)](https://github.com/no-src/log/releases)

## Installation

`go get -u github.com/no-src/log`

## Quick Start

Current support following loggers

- empty logger
- console logger
- file logger
- multi logger

For example, init a file logger, to write logs.

```go
log.InitDefaultLogger(NewFileLogger(DebugLevel,"./logs",""))
log.Debug("%s,test debug log", "hello")
log.Info("%s,test info log", "hello")
log.Warn("%s,test warn log", "hello")
log.Error(errors.New("log err"), "%s,test error log", "hello")
log.Log("%s,test log log", "hello")
```