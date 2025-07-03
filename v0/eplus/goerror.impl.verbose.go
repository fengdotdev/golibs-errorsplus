package eplus

var _ Verbose = (*GoError)(nil)

// VerboseError implements Verbose.
func (e *GoError) VerboseError() string {
	panic("unimplemented")
}
