package json

import (
	"bytes"
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
	buf := new(bytes.Buffer)
	encoder := json.NewEncoder(buf)
	encoder.SetEscapeHTML(false)
	// Encode will auto append row terminator
	if err := encoder.Encode(c); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func init() {
	formatter.Register(formatter.JsonFormatter, newJsonFormatter())
}
