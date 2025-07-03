package eplus

func New(err error) ErrorPlus {
	if err == nil {
		return nil
	}

	traceSlice := trace()
	tags := make([]string, 0)
	goerr := NewGoError("", err, traceSlice, tags)
	return goerr
}

func NewError(err error, msg string, tags []string) ErrorPlus {
	return nil
}
