package content

import (
	"time"

	"github.com/no-src/log/level"
)

// Content the log content info
type Content struct {
	Level      level.Level   `json:"level"`
	Time       Time          `json:"time"`
	Log        string        `json:"log"`
	Error      error         `json:"-"`
	AppendTime bool          `json:"-"`
	Args       []interface{} `json:"-"`
}

// NewContent return an instance of Content
func NewContent(lvl level.Level, log string, args []interface{}, err error, appendTime bool) Content {
	c := Content{
		Level:      lvl,
		Time:       Time(time.Now()),
		Log:        log,
		Args:       args,
		AppendTime: appendTime,
	}
	if err != nil {
		c.Error = Error{
			err: err,
		}
	}
	return c
}
