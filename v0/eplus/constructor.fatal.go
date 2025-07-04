package eplus

func Fatal(err error) ErrorPlus {
	if err == nil {
		return nil
	}

	traceStack := trace()
	tags := make([]string, 0)
	msg := err.Error()

	goerr := NewGoError(msg, err, traceStack, tags)
	goerr.severity = FatalSev
	return goerr
}
