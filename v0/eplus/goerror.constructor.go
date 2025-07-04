package eplus

// default constructor for GoError
func NewGoError(msg string, selfError error, trace []Trace, tags []string) *GoError {
	if msg == "" && selfError == nil {
		return nil
	}

	if selfError != nil && msg == "" {
		msg = selfError.Error()
	}

	settagsMap := make(map[string]struct{}, len(tags))
	for _, tag := range tags {
		settagsMap[tag] = struct{}{}
	}

	return &GoError{
		child:       nil,
		msg:         msg,
		selfError:   selfError,
		trace:       trace,
		tags:        settagsMap,
		severity:    UnknownSev,   // default severity
		known:       UnknownError, // default is UnknownError
		temporality: Permanent,    // default is Permanent
	}
}

func WrapGoError(wrap error, newerr error, msg string, trace []Trace, tags []string) *GoError {

	ep := NewGoError(msg, newerr, trace, tags)

	if ep == nil {
		return nil
	}
	ep.child = wrap

	return ep
}
