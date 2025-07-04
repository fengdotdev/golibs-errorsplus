package eplus

var _ ErrorPlus = (*GoError)(nil)

type GoError struct {
	child     error
	selfError error
	msg       string
	trace     []Trace
	tags      map[string]struct{} // set of tags
}
