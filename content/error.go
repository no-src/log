package content

// Error the custom error for log
type Error struct {
	err error
}

// NewError return a custom error wrap the real error
func NewError(err error) error {
	return Error{
		err: err,
	}
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
