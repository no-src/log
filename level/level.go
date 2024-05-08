package level

import (
	"errors"
	"strings"
)

// Level the log level
type Level int8

const (
	// unknownLevel the unknown log level
	unknownLevel Level = iota - 1
	// DebugLevel the debug log level
	DebugLevel
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

// MarshalText implement interface encoding.TextMarshaler
func (l Level) MarshalText() (text []byte, err error) {
	return []byte(l.String()), nil
}

// ParseLevel parse the log level from string
func ParseLevel(text string) (Level, error) {
	text = strings.ToUpper(text)
	switch text {
	case DebugLevel.String():
		return DebugLevel, nil
	case InfoLevel.String():
		return InfoLevel, nil
	case WarnLevel.String():
		return WarnLevel, nil
	case ErrorLevel.String():
		return ErrorLevel, nil
	case NoneLevel.String():
		return NoneLevel, nil
	default:
		return unknownLevel, errors.New("invalid level")
	}
}
