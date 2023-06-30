package content

import (
	"time"

	"github.com/no-src/log/level"
)

// Content the log content info
type Content struct {
	Level      level.Level `json:"level"`
	Time       *Time       `json:"time,omitempty"`
	Log        string      `json:"log"`
	Error      error       `json:"-"`
	AppendTime bool        `json:"-"`
	Args       []any       `json:"-"`
}

// NewContent return an instance of Content
func NewContent(lvl level.Level, err error, appendTime bool, timeFormat string, log string, args ...any) Content {
	var t *Time
	if appendTime {
		t = NewTimeWithFormat(time.Now(), timeFormat)
	}
	return NewContentWithTime(lvl, err, t, log, args...)
}

// NewContentWithTime return an instance of Content with specified time
func NewContentWithTime(lvl level.Level, err error, t *Time, log string, args ...any) Content {
	c := Content{
		Level:      lvl,
		Log:        log,
		Args:       args,
		AppendTime: t != nil,
		Time:       t,
	}
	if err != nil {
		c.Error = NewError(err)
	}
	return c
}
