package content

import (
	"testing"
	"time"
)

func TestTime_MarshalText(t *testing.T) {
	now := time.Now()
	expect := now.Format(DefaultLogTimeFormat)
	ti := NewTime(time.Now())
	data, err := ti.MarshalText()
	if err != nil {
		t.Errorf("Time.MarshalText error => %v", err)
		return
	}
	actual := string(data)
	if expect != actual {
		t.Errorf("test Time.MarshalText failed, expect to get %s, but actual get %s", expect, actual)
	}
}
