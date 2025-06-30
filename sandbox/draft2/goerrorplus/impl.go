package goerrorplus

import (
	"errors"
	"fmt"
)

func (g *GoErrorPlus) VerboseError() error {
	timestr := g.time.Format(TimeFormat)
	trace := g.trace
	errstr := g.err.Error()

	s := fmt.Sprintf("time: %s | trace: %s | message: %s | error: %s", timestr, trace, g.message, errstr)

	// example: go ver.: go1.21.1 | at: 2023-10-03 16:13:19 | msg: recover: div by zero | error: div failed | tags: [math-err] | fn: main.division | trace: [caller: G:/fengdotdev/github/goerrorsplus/example/main.go line: 31] | args: [%!s(int=1) %!s(int=0)]
	return errors.New(s)

}
