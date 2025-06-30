package goerrorplus

import (
	"fmt"
	"runtime"
	"time"
)

func New(err error) *GoErrorPlus {
	if err == nil {
		return nil
	}

	_, caller, line, _ := runtime.Caller(3)

	return &GoErrorPlus{
		err:          err,
		trace:        fmt.Sprintf("[caller: %s line: %d]", caller, line),
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
