package eplus

func NewGoError(msg string, selfError error, trace []string, tags []string) *GoError {
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
		child:     nil,
		msg:       msg,
		selfError: selfError,
		trace:     trace,
		tags:      settagsMap,
	}
}
