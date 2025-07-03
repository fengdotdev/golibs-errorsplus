package errorplus

var _ error = (*GoError)(nil)

// todo check behavior of nil error
func (e *GoError) Error() string {
	if e == nil {
		return "nil error"
	}
	if e.selfError != nil {
		return e.selfError.Error()
	}
	return e.msg
}
