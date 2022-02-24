package log

// Level the log level
type Level int8

const (
	// DebugLevel the debug log level
	DebugLevel Level = iota
	// InfoLevel the info log level
	InfoLevel
	// WarnLevel the warn log level
	WarnLevel
	// ErrorLevel the error log level
	ErrorLevel
	// NoneLevel disable all level log, except for the Writer.Log
	NoneLevel
)

// String return the description of log level
func (l Level) String() string {
	switch l {
	case DebugLevel:
		return "DEBUG"
	case InfoLevel:
		return "INFO"
	case WarnLevel:
		return "WARN"
	case ErrorLevel:
		return "ERROR"
	case NoneLevel:
		return "NONE"
	}
	return "UNKNOWN"
}
