package json

import (
	"encoding/json"
	"fmt"

	"github.com/no-src/log/content"
	"github.com/no-src/log/formatter"
)

type jsonFormatter struct {
}

func newJsonFormatter() formatter.Formatter {
	return &jsonFormatter{}
}

func (f *jsonFormatter) Serialize(c content.Content) ([]byte, error) {
	c.Log = fmt.Sprintf(c.Log, c.Args...)
	if c.Error != nil {
		c.Log = fmt.Sprintf(c.Log+". %s", c.Error)
	}
	data, err := json.Marshal(c)
	if err != nil {
		return nil, err
	}
	data = append(data, []byte(formatter.AppendRowTerminator(""))...)
	return data, err
}

func init() {
	formatter.Register(formatter.JsonFormatter, newJsonFormatter())
}
