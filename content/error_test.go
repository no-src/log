package content

import (
	"io"
	"testing"
)

func TestError_MarshalText(t *testing.T) {
	testCases := []struct {
		name   string
		err    error
		expect string
	}{
		{"with error", io.EOF, "EOF"},
		{"with nil error", nil, ""},
	}
	for _, tc := range testCases {
		cErr := NewError(tc.err).(Error)
		data, err := cErr.MarshalText()
		if err != nil {
			t.Errorf("Error.MarshalText error => %v", err)
			return
		}
		actual := string(data)
		if tc.expect != actual {
			t.Errorf("test Error.MarshalText failed, expect to get %s, but actual get %s", tc.expect, actual)
		}
	}
}
