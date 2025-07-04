package eplus

var _ Verbose = (*GoError)(nil)

// VerboseError implements Verbose.
func (e *GoError) VerboseError() string {

	ftmtrace := ""

	if trace == nil {
		ftmtrace = ""
	} else {
		for _, f := range e.trace {

			result := f.Function + " at " + f.File + ":" + string(f.Line) + "\n"

			ftmtrace += result
		}
	}

	return ftmtrace
}
