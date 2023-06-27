package json

import (
	"errors"
	"testing"
	"time"

	"github.com/no-src/log/content"
	"github.com/no-src/log/level"
)

func TestJsonFormatter_Serialize(t *testing.T) {
	logTime, _ := time.ParseInLocation("2006-01-02 15:04:05", "2022-06-25 23:59:59", time.UTC)
	logTimeP := content.NewTime(logTime)

	testCases := []struct {
		name    string
		content content.Content
		expect  string
	}{
		{"debug no args", content.NewContentWithTime(level.DebugLevel, nil, nil, "<json formatter> hello"), `{"level":"DEBUG","log":"<json formatter> hello"}`},
		{"debug with args", content.NewContentWithTime(level.DebugLevel, nil, nil, "<json formatter> %s %s", "hello", "world"), `{"level":"DEBUG","log":"<json formatter> hello world"}`},
		{"error no args", content.NewContentWithTime(level.DebugLevel, errors.New("test error"), nil, "<json formatter> hello"), `{"level":"DEBUG","log":"<json formatter> hello. test error"}`},
		{"error with args", content.NewContentWithTime(level.DebugLevel, errors.New("test error"), nil, "<json formatter> %s %s", "hello", "world"), `{"level":"DEBUG","log":"<json formatter> hello world. test error"}`},

		{"debug no args and append time", content.NewContentWithTime(level.DebugLevel, nil, logTimeP, "<json formatter> hello"), `{"level":"DEBUG","time":"2022-06-25 23:59:59","log":"<json formatter> hello"}`},
		{"debug with args and append time", content.NewContentWithTime(level.DebugLevel, nil, logTimeP, "<json formatter> %s %s", "hello", "world"), `{"level":"DEBUG","time":"2022-06-25 23:59:59","log":"<json formatter> hello world"}`},
		{"error no args and append time", content.NewContentWithTime(level.DebugLevel, errors.New("test error"), logTimeP, "<json formatter> hello"), `{"level":"DEBUG","time":"2022-06-25 23:59:59","log":"<json formatter> hello. test error"}`},
		{"error with args and append time", content.NewContentWithTime(level.DebugLevel, errors.New("test error"), logTimeP, "<json formatter> %s %s", "hello", "world"), `{"level":"DEBUG","time":"2022-06-25 23:59:59","log":"<json formatter> hello world. test error"}`},
		{"error with args that contain the format character", content.NewContentWithTime(level.DebugLevel, errors.New("test error"), logTimeP, "<json formatter> %s %s", "hello%s%v", "world"), `{"level":"DEBUG","time":"2022-06-25 23:59:59","log":"<json formatter> hello%s%v world. test error"}`},
	}
	f := newJsonFormatter()
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			data, err := f.Serialize(tc.content)
			if err != nil {
				t.Errorf("test Serialize for json formatter error => %s", err)
				return
			}
			tc.expect = tc.expect + "\n"
			actual := string(data)
			if tc.expect != actual {
				t.Errorf("test Serialize for json formatter error, expect: %s, but actual: %s", tc.expect, actual)
			}
		})
	}
}
