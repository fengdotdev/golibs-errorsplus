package eplus

var _ ErrorPlusTrace = (*GoError)(nil)

// Moficated Striped first and last
func (e *GoError) Trace() (output []string) {
	if e == nil {
		return nil
	}

	if e.trace == nil {
		return nil
	}

	output = make([]string, 0, len(e.trace))

	// format the trace removing runtime-callers and the package errorplus
	fmttrace := fmttrace(e.trace)
	if len(fmttrace) == 0 {
		output = append(output, "No trace available")
		return output
	}

	for _, t := range e.trace {
		output = append(output, t.String())
	}

	return output
}

// Moficated Striped first and last
func (e *GoError) TraceString() (output string) {

	defer func() {
		if r := recover(); r != nil {
			// if panic, return empty string
			output = ""
		}
	}()

	if e.trace != nil {

		startArt := "	-----> "
		At := " at: "
		Line := " line: "
		Func := "func: "
		endArt := "\n"

		// format the trace removing runtime-callers and the package errorplus
		fmttrace := fmttrace(e.trace)

		if len(fmttrace) == 0 {
			output = "No trace available\n"
		} else {
			output = "Trace:\n"

			for _, f := range fmttrace {

				result := startArt + Func + f.Function + At + f.File + Line + f.LineStr() + endArt

				output += result
			}
		}

	}
	return output
}

func (e *GoError) TraceRaw() []Trace {
	if e == nil {
		return nil
	}
	return e.trace
}

func (e *GoError) TraceFull() []string {
	if e == nil {
		return nil
	}

	if e.trace == nil {
		return nil
	}

	output := make([]string, 0, len(e.trace))

	for _, t := range e.trace {
		output = append(output, t.String())
	}

	return output
}
