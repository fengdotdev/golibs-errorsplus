package errorplus

import (
	"fmt"
	"runtime"
)

func trace() []string {

	trace := []string{}

	// := 2 // Skip the current function and the caller of this function
	frames := runtime.CallersFrames(make([]uintptr, 0, 64))

	//n := runtime.Callers(skip, nil)

	for {
		frame, more := frames.Next()
		AreWeFinish := !more
		result := fmt.Sprintf("FUNC: %s AT: %s:%d\n", frame.Function, frame.File, frame.Line)

		if frame.Function == packageName || frame.Function == packageName {
			// do nothing,
		} else {
			trace = append(trace, result)
		}

		if AreWeFinish {
			break
		}

	}
	return trace
}
