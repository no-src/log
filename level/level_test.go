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
