package content

import (
	"io"
	"testing"
)

func TestError_MarshalText(t *testing.T) {
	cErr := NewError(io.EOF).(Error)
	data, err := cErr.MarshalText()
	if err != nil {
		t.Errorf("Error.MarshalText error => %v", err)
		return
	}
	expect := io.EOF.Error()
	actual := string(data)
	if expect != actual {
		t.Errorf("test Error.MarshalText failed, expect to get %s, but actual get %s", expect, actual)
	}
}
