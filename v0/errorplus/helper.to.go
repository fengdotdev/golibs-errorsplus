package errorplus

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
