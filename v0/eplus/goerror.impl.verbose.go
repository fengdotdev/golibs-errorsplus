package eplus

var _ Verbose = (*GoError)(nil)

// VerboseError implements Verbose.
func (e *GoError) VerboseError() string {

	output := ""

	if e.trace == nil {
		output = ""
	} else {

		// format the trace
		fmttrace := fmttrace(e.trace)

		if len(fmttrace) == 0 {
			output = "No trace available\n"
		} else {
			output = "Trace:\n"

			for _, f := range fmttrace {

				result := "-----> " + "func: " + f.Function + " at: " + f.File + " line: " + f.LineStr() + "\n"

				output += result
			}
		}

	}

	return output
}
