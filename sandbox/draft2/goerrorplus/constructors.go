package goerrorplus

import (
	"fmt"
	"reflect"
	"runtime"
	"time"
)

func New(err error) *GoErrorPlus {
	if err == nil {
		return nil
	}
	pcs := make([]uintptr, 15)
	n := runtime.Callers(4, pcs) // saltando runtime.Callers y logStackTrace
	frames := runtime.CallersFrames(pcs[:n-2])

	callerfn := runtime.FuncForPC(reflect.ValueOf(New.Pointer()).Name())

	out := "\n"
	for {
		frame, more := frames.Next()
		result := fmt.Sprintf("FUNC: %s AT: %s:%d\n", frame.Function, frame.File, frame.Line)

		if frame.Function == "tuPaquete.New" || frame.Function == "tuPaquete.logError" {
			continue
		}

		out += result
		if !more {
			break
		}
	}
	return &GoErrorPlus{
		err:          err,
		trace:        out,
		time:         time.Now(),
		runtimeGoVer: runtime.Version(),
	}
}

func NewError(err error, message string, tags []string) *GoErrorPlus {

	if err == nil {
		return nil
	}

	g := New(err)
	g.message = message
	g.tags = tags

	return g
}

func NewWithMessage(err error, message string) *GoErrorPlus {
	if err == nil {
		return nil
	}

	g := New(err)
	g.message = message
	return g
}
func NewWithTags(err error, tags ...string) *GoErrorPlus {
	if err == nil {
		return nil
	}

	g := New(err)
	g.tags = tags
	return g
}

func NewWithOptions(err error, op Options) *GoErrorPlus {
	if err == nil {
		return nil
	}

	g := New(err)

	return g
}
