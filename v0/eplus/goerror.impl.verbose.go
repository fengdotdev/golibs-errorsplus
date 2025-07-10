package eplus

import "github.com/fengdotdev/golibs-errorsplus/v0/customtypes/try"

var _ Verbose = (*GoError)(nil)

// VerboseError implements Verbose.
func (e *GoError) Verbose() string {

	output := ""

	try.If(e == nil).Try(func() {
		output = "GoError is nil"
		return
	})

	try.If(showMessage.Get()).Try(func() {
		output = e.msg + "\n"
	})

	try.If(showTrace.Get()).Try(func() {
		output += e.TraceString() + "\n"
	})

	return output
}
