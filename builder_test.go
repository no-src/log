package log

import (
	"testing"
)

func TestCreateLoggerFromConfig(t *testing.T) {
	logger, err := CreateLoggerFromConfig("./testdata/conf.yaml")
	if err != nil {
		t.Fatal(err)
	}
	InitDefaultLogger(logger)
	defer Close()
	testLogs(t)
}
