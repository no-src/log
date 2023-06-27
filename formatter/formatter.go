package formatter

import (
	"strings"
	"sync"

	"github.com/no-src/log/content"
)

// Formatter the log formatter interface
type Formatter interface {
	// Serialize serialize the log content to []byte
	Serialize(c content.Content) ([]byte, error)
}

var (
	formatters           = make(map[string]Formatter)
	defaultFormatterType = TextFormatter
	defaultTerminator    = "\n"
	mu                   sync.RWMutex
)

// Default return the global default Formatter
func Default() Formatter {
	mu.RLock()
	defer mu.RUnlock()
	return New(defaultFormatterType)
}

// NewJsonFormatter return a json formatter
func NewJsonFormatter() Formatter {
	return New(JsonFormatter)
}

// NewTextFormatter return a text formatter
func NewTextFormatter() Formatter {
	return New(TextFormatter)
}

// InitDefaultFormatter init the global default Formatter by specified type
func InitDefaultFormatter(t string) {
	_, ok := formatters[t]
	if ok {
		mu.Lock()
		defaultFormatterType = t
		mu.Unlock()
	}
}

// New return a Formatter by specified type.
// if the specified Formatter does not exist and return the default Formatter.
func New(t string) Formatter {
	f := formatters[t]
	if f == nil {
		return Default()
	}
	return f
}

// Register register a Formatter
func Register(t string, formatter Formatter) {
	formatters[t] = formatter
}

// AppendRowTerminator append a terminator at the end of the row if that does not exist
func AppendRowTerminator(format string) string {
	if !strings.HasSuffix(format, defaultTerminator) {
		format = format + defaultTerminator
	}
	return format
}
