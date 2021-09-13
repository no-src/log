# log

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