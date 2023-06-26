package log

import (
	"testing"

	"github.com/no-src/log/formatter"
)

func TestEmptyLogger(t *testing.T) {
	testCases := []struct {
		name      string
		formatter string
	}{
		{"TextFormatter", formatter.TextFormatter},
		{"JsonFormatter", formatter.JsonFormatter},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			InitDefaultLogger(NewEmptyLogger().WithFormatter(formatter.New(tc.formatter)).WithTimeFormat(testTimeFormat))
			defer Close()
			testLogs(t)
		})
	}
}
