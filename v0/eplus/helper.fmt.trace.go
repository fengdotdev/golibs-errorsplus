package eplus

import "github.com/fengdotdev/golibs-errorsplus/v0/customtypes/ontry"

func fmttrace(traceIn []Trace) (traceOut []Trace) {

	defer func() {
		if r := recover(); r != nil {
			// If we panic, we return an empty trace slice
			traceOut = []Trace{}
			return
		}
	}()

	traceOut = []Trace{}

	if traceIn == nil {
		return traceOut
	}
	nskip := traceSkipFirst.Get()

	nlast := traceSkipLast.Get()

	ontry.If(len(traceIn) > nskip).Try(func() {
		// remove first elements as defined in traceSkipFirst
		traceIn = traceIn[nskip:]
	})

	ontry.If(len(traceIn) > nlast).Try(func() {
		// remove last elements as defined in traceSkipLast
		traceIn = traceIn[:len(traceIn)-nlast]
	})

	for _, f := range traceIn {
		result := Trace{
			File:     f.File,
			Function: f.Function,
			Line:     f.Line,
		}
		traceOut = append(traceOut, result)
	}

	return traceOut
}
