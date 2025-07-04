package eplus

import (
	"fmt"
	"runtime"
)

type Trace struct {
	File     string
	Function string
	Line     string
}

func trace() []Trace {

	traceStack := []Trace{}

	const maxDepth = 15
	pcs := make([]uintptr, maxDepth)

	// skip=3:
	// 0 = runtime.Callers
	// 1 = New
	// 2 = función que llama a New
	skip := 3

	n := runtime.Callers(skip, pcs)
	frames := runtime.CallersFrames(pcs[:n])

	for {
		frame, more := frames.Next()
		AreWeFinish := !more

		//if frame.Function == "runtime.goexit" {
		//	continue
		//}

		trace := Trace{
			File:     frame.File,
			Function: frame.Function,
			Line:     fmt.Sprintf("%d", frame.Line),
		}
		traceStack = append(traceStack, trace)

		if AreWeFinish {
			break
		}

	}

	// remove last 2 elements for runtime.main runtime.goexit

	if len(traceStack) > 2 {
		traceStack = traceStack[:len(traceStack)-2]
	}
	return traceStack
}
