package log

import "testing"

func TestLevel(t *testing.T) {
	var lvl Level = 0
	if lvl.String() != "DEBUG" {
		t.Errorf("get debug level error")
		return
	}
	lvl = 1
	if lvl.String() != "INFO" {
		t.Errorf("get info level error")
		return
	}

	lvl = 2
	if lvl.String() != "WARN" {
		t.Errorf("get warn level error")
		return
	}

	lvl = 3
	if lvl.String() != "ERROR" {
		t.Errorf("get error level error")
		return
	}

	lvl = 4
	if lvl.String() != "NONE" {
		t.Errorf("get none level error")
		return
	}

	lvl = 5
	if lvl.String() != "UNKNOWN" {
		t.Errorf("get unknown level error")
		return
	}

	lvl = -1
	if lvl.String() != "UNKNOWN" {
		t.Errorf("get unknown level error")
		return
	}
}
