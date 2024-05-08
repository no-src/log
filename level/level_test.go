package level

import (
	"testing"
)

func TestLevel(t *testing.T) {
	testCases := []struct {
		lvl    int
		expect string
	}{
		{0, "DEBUG"},
		{1, "INFO"},
		{2, "WARN"},
		{3, "ERROR"},
		{4, "NONE"},
		{5, "UNKNOWN"},
		{-1, "UNKNOWN"},
	}

	for _, tc := range testCases {
		t.Run(tc.expect, func(t *testing.T) {
			actual := Level(tc.lvl).String()
			if actual != tc.expect {
				t.Errorf("get log level error, [%d] expect:%s, actual:%s", tc.lvl, tc.expect, actual)
				return
			}
			marshalData, err := Level(tc.lvl).MarshalText()
			if err != nil {
				t.Errorf("marshal log level error, [%d] => %s", tc.lvl, err)
				return
			}
			if string(marshalData) != tc.expect {
				t.Errorf("marshal log level error, [%d] expect:%s, actual:%s", tc.lvl, tc.expect, string(marshalData))
			}
		})
	}
}

func TestParseLevel(t *testing.T) {
	testCases := []struct {
		expect Level
		text   string
	}{
		{DebugLevel, "DEBUG"},
		{InfoLevel, "INFO"},
		{WarnLevel, "WARN"},
		{ErrorLevel, "ERROR"},
		{NoneLevel, "NONE"},

		{DebugLevel, "debug"},
		{InfoLevel, "info"},
		{WarnLevel, "warn"},
		{ErrorLevel, "error"},
		{NoneLevel, "none"},
	}

	for _, tc := range testCases {
		t.Run(tc.text, func(t *testing.T) {
			actual, err := ParseLevel(tc.text)
			if err != nil {
				t.Errorf("ParseLevel error, text=%s err=%s", tc.text, err)
				return
			}
			if actual != tc.expect {
				t.Errorf("ParseLevel error, [%s] expect:%d, actual:%d", tc.text, tc.expect, actual)
			}
		})
	}
}

func TestParseLevel_ReturnError(t *testing.T) {
	testCases := []struct {
		expect Level
		text   string
	}{
		{unknownLevel, "UNKNOWN"},
		{unknownLevel, "INVALID"},
		{unknownLevel, "HELLO"},

		{unknownLevel, "unknown"},
		{unknownLevel, "invalid"},
		{unknownLevel, "hello"},
		{unknownLevel, ""},
	}

	for _, tc := range testCases {
		t.Run(tc.text, func(t *testing.T) {
			actual, err := ParseLevel(tc.text)
			if err == nil {
				t.Errorf("ParseLevel expect to get an error but get nil, text=%s", tc.text)
				return
			}
			if actual != tc.expect {
				t.Errorf("ParseLevel error, [%s] expect:%d, actual:%d", tc.text, tc.expect, actual)
			}
		})
	}
}
