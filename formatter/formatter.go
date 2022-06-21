package formatter

import (
	"strings"

	"github.com/no-src/log/content"
)

// Formatter the log formatter interface
type Formatter interface {
	// Serialize serialize the log content to []byte
	Serialize(c content.Content) ([]byte, error)
}

var (
	formatters           = make(map[Type]Formatter)
	defaultFormatterType = TextFormatter
	defaultTerminator    = "\n"
)

// Default return the global default Formatter
func Default() Formatter {
	return New(defaultFormatterType)
}

// InitDefaultFormatter init the global default Formatter by specified type
func InitDefaultFormatter(t Type) {
	_, ok := formatters[t]
	if ok {
		defaultFormatterType = t
	}
}

// New return a Formatter by specified type
func New(t Type) Formatter {
	f := formatters[t]
	if f == nil {
		return Default()
	}
	return f
}

// Register register a Formatter
func Register(t Type, formatter Formatter) {
	formatters[t] = formatter
}

// AppendRowTerminator append a terminator at the end of the row if that does not exist
func AppendRowTerminator(format string) string {
	if !strings.HasSuffix(format, defaultTerminator) {
		format = format + defaultTerminator
	}
	return format
}
