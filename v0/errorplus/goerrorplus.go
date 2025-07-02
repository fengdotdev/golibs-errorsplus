package errorplus

type GoError struct {
	child     error
	selfError error
	msg       string
	trace     []string
}

// todo check behavior of nil error
func (e *GoError) Error() string {
	if e == nil {
		return "nil error"
	}
	if e.selfError != nil {
		return e.selfError.Error()
	}
	return e.msg
}

func ToGoError(err error) *GoError {
	if err == nil {
		return &GoError{msg: "", trace: []string{}} // GoError vacío para nil
	}

	if gerr, ok := err.(*GoError); ok {
		return gerr
	}

	return &GoError{
		msg:   err.Error(),
		trace: []string{},
	}
}
