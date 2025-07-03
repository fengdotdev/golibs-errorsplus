package eplus

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

func ToErrorPlus(err error) ErrorPlus {
	if err == nil {
		return nil
	}

	if eplusErr, ok := err.(ErrorPlus); ok {
		return eplusErr
	}

	goErr := ToGoError(err)
	return New(goErr)
}
