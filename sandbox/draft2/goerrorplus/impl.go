package goerrorplus

import (
	"errors"
	"fmt"
)

func (g *GoErrorPlus) VerboseError() error {

	timestamp := g.time.Format(TimeFormat)
	version := g.runtimeGoVer
	err := g.err.Error()
	message := g.message
	tags := g.tags
	trace := g.trace
	fn := g.fN
	args := g.args

	s := fmt.Sprintf("go ver.: %s | at: %s | msg: %s | error: %s | tags: %s | fn: %s | trace: %s | args: %s", version, timestamp, err, message, tags, fn, trace, args)

	// example: go ver.: go1.21.1 | at: 2023-10-03 16:13:19 | msg: recover: div by zero | error: div failed | tags: [math-err] | fn: main.division | trace: [caller: G:/fengdotdev/github/goerrorsplus/example/main.go line: 31] | args: [%!s(int=1) %!s(int=0)]
	return errors.New(s)

}
