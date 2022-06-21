package content

// Error the custom error for log
type Error struct {
	err error
}

// Error implement interface error
func (e Error) Error() string {
	if e.err == nil {
		return ""
	}
	return e.err.Error()
}

// MarshalText implement interface encoding.TextMarshaler
func (e Error) MarshalText() (text []byte, err error) {
	return []byte(e.Error()), nil
}
