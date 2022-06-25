package content

import (
	"time"

	"github.com/no-src/log/level"
)

// Content the log content info
type Content struct {
	Level      level.Level   `json:"level"`
	Time       *Time         `json:"time,omitempty"`
	Log        string        `json:"log"`
	Error      error         `json:"-"`
	AppendTime bool          `json:"-"`
	Args       []interface{} `json:"-"`
}

// NewContent return an instance of Content
func NewContent(lvl level.Level, err error, appendTime bool, log string, args ...interface{}) Content {
	c := Content{
		Level:      lvl,
		Log:        log,
		Args:       args,
		AppendTime: appendTime,
	}
	if appendTime {
		c.Time = NewTime(time.Now())
	}
	if err != nil {
		c.Error = Error{
			err: err,
		}
	}
	return c
}
