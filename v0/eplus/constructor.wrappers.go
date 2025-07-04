package eplus

import "errors"

func Wrap(wrap error, err error) ErrorPlus {
	if err == nil {
		return nil
	}

	traceStack := trace()
	tags := make([]string, 0)
	msg := err.Error()

	goerr := WrapGoError(wrap, err, msg, traceStack, tags)
	return goerr
}

func WrapError(wrap error, err string) ErrorPlus {
	if err == "" {
		return nil
	}

	traceStack := trace()
	tags := make([]string, 0)
	msg := err

	goerr := WrapGoError(wrap, errors.New(err), msg, traceStack, tags)
	return goerr
}
func WrapErrorWithTags(wrap error, err string, tags []string) ErrorPlus {
	if err == "" {
		return nil
	}

	traceStack := trace()
	msg := err

	goerr := WrapGoError(wrap, errors.New(err), msg, traceStack, tags)
	return goerr
}
