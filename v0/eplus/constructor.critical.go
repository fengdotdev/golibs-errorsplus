package eplus

import "errors"

func Critical(err error) ErrorPlus {
	if err == nil {
		return nil
	}

	traceStack := trace()
	tags := make([]string, 0)
	msg := err.Error()

	goerr := NewGoError(msg, err, traceStack, tags)
	goerr.severity = CriticalSev
	return goerr
}

func CriticalError(err string) ErrorPlus {
	if err == "" {
		return nil
	}

	traceStack := trace()
	tags := make([]string, 0)
	msg := err

	goerr := NewGoError(msg, errors.New(err), traceStack, tags)
	goerr.severity = CriticalSev
	return goerr
}
