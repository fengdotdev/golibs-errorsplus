package eplus

var _ ErrorPlus = (*GoError)(nil)

type GoError struct {
	child       error               // the error that is wrapped by this error if exists any
	selfError   error               // the error that is passed to the constructor
	msg         string              // if empty, use selfError.Error()
	trace       []Trace             // stack trace of the error if available with all frames
	tags        map[string]struct{} // set of tags
	severity    Severity            // default fatal
	known       Knowledge           // default is UnknownError
	temporality Temporality         // default is Permanent
	
}
