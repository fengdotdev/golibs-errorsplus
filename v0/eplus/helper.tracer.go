package eplus

import (
	"fmt"
	"runtime"
)

// trace returns a slice of Trace structs representing the call stack (all of it)
func trace() (traceStack []Trace) {
	traceStack = []Trace{}

	// this func may fail if the stack is too deep or other runtime issues,
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in trace:", r)
			// return empty trace stack
			traceStack = []Trace{}
			return
		}
	}()

	maxDepth := traceMaxDepth.Get()
	pcs := make([]uintptr, maxDepth)

	n := runtime.Callers(0, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()
		AreWeFinish := !more

		trace := Trace{
			File:     frame.File,
			Function: frame.Function,
			Line:     frame.Line,
		}
		traceStack = append(traceStack, trace)

		if AreWeFinish {
			break
		}

	}

	return traceStack
}
