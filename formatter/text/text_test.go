package text

import (
	"errors"
	"testing"
	"time"

	"github.com/no-src/log/content"
	"github.com/no-src/log/level"
)

func TestTextFormatter_Serialize(t *testing.T) {
	logTime, _ := time.ParseInLocation(content.DefaultLogTimeFormat, "2022-06-25 23:59:59", time.UTC)
	logTimeP := content.NewTime(logTime)

	testCases := []struct {
		name    string
		content content.Content
		expect  string
	}{
		{"debug no args", content.NewContentWithTime(level.DebugLevel, nil, nil, "<text formatter> hello"), "[DEBUG] <text formatter> hello"},
		{"debug with args", content.NewContentWithTime(level.DebugLevel, nil, nil, "<text formatter> %s %s", "hello", "world"), "[DEBUG] <text formatter> hello world"},
		{"error no args", content.NewContentWithTime(level.DebugLevel, errors.New("test error"), nil, "<text formatter> hello"), "[DEBUG] <text formatter> hello. test error"},
		{"error with args", content.NewContentWithTime(level.DebugLevel, errors.New("test error"), nil, "<text formatter> %s %s", "hello", "world"), "[DEBUG] <text formatter> hello world. test error"},

		{"debug no args and append time", content.NewContentWithTime(level.DebugLevel, nil, logTimeP, "<text formatter> hello"), "[2022-06-25 23:59:59] [DEBUG] <text formatter> hello"},
		{"debug with args and append time", content.NewContentWithTime(level.DebugLevel, nil, logTimeP, "<text formatter> %s %s", "hello", "world"), "[2022-06-25 23:59:59] [DEBUG] <text formatter> hello world"},
		{"error no args and append time", content.NewContentWithTime(level.DebugLevel, errors.New("test error"), logTimeP, "<text formatter> hello"), "[2022-06-25 23:59:59] [DEBUG] <text formatter> hello. test error"},
		{"error with args and append time", content.NewContentWithTime(level.DebugLevel, errors.New("test error"), logTimeP, "<text formatter> %s %s", "hello", "world"), "[2022-06-25 23:59:59] [DEBUG] <text formatter> hello world. test error"},
		{"error with args that contain the format character", content.NewContentWithTime(level.DebugLevel, errors.New("test error"), logTimeP, "<text formatter> %s %s", "hello%s%v", "world"), "[2022-06-25 23:59:59] [DEBUG] <text formatter> hello%s%v world. test error"},
	}
	f := newTextFormatter()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := f.Serialize(tc.content)
			if err != nil {
				t.Errorf("test Serialize for text formatter error => %s", err)
				return
			}
			tc.expect = tc.expect + "\n"
			actual := string(data)
			if tc.expect != actual {
				t.Errorf("test Serialize for text formatter error, expect: %s, but actual: %s", tc.expect, actual)
			}
		})
	}
}
