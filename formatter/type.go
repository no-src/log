package formatter

// Type the log formatter type
type Type int

const (
	// UnknownFormatter the unknown formatter type
	UnknownFormatter Type = iota
	// TextFormatter the text formatter type
	TextFormatter
	// JsonFormatter the json formatter type
	JsonFormatter
)
