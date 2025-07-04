package eplus

import "errors"

func New(err error) ErrorPlus {
	if err == nil {
		return nil
	}

	traceStack := trace()
	tags := make([]string, 0)
	msg := err.Error()
	goerr := NewGoError(msg, err, traceStack, tags)
	return goerr
}

func NewWithTags(err error, tags []string) ErrorPlus {
	if err == nil {
		return nil
	}

	traceStack := trace()
	msg := err.Error()
	goerr := NewGoError(msg, err, traceStack, tags)
	return goerr
}

func NewWithTagsAndMessage(err error, msg string, tags []string) ErrorPlus {
	if err == nil {
		return nil
	}

	traceStack := trace()
	goerr := NewGoError(msg, err, traceStack, tags)
	return goerr
}

func NewError(err string) ErrorPlus {

	traceStack := trace()
	tags := make([]string, 0)
	msg := err
	goerr := NewGoError(msg, errors.New(err), traceStack, tags)
	return goerr
}

func NewErrorWithTags(err string, tags []string) ErrorPlus {
	if err == "" {
		return nil
	}

	traceStack := trace()
	msg := err
	goerr := NewGoError(msg, errors.New(err), traceStack, tags)
	return goerr
}

