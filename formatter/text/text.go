package text

import (
	"fmt"
	"github.com/no-src/log/content"
	"github.com/no-src/log/formatter"
)

type textFormatter struct {
}

func newTextFormatter() formatter.Formatter {
	return &textFormatter{}
}

func (f *textFormatter) Serialize(c content.Content) ([]byte, error) {
	var format = "%s[%s] %s" // [time] [level] content<. error>
	var timeSection string
	if c.AppendTime && c.Time != nil {
		timeSection = fmt.Sprintf("[%s] ", c.Time.Time().Format(content.LogTimeFormat))
	}
	content := fmt.Sprintf(c.Log, c.Args...)
	format = fmt.Sprintf(format, timeSection, c.Level.String(), content)
	if c.Error != nil {
		format = fmt.Sprintf(format+". %s", c.Error)
	}
	format = formatter.AppendRowTerminator(format)
	return []byte(format), nil
}

func init() {
	formatter.Register(formatter.TextFormatter, newTextFormatter())
}
