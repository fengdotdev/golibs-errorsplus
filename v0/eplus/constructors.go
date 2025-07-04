package eplus

func New(err error) ErrorPlus {
	if err == nil {
		return nil
	}

	traceStack := trace()
	tags := make([]string, 0)
	goerr := NewGoError("", err, traceStack, tags)
	return goerr
}

func NewError(err error, msg string, tags []string) ErrorPlus {
	return nil
}
