package goerrorplus

import "github.com/fengdotdev/golibs-errorsplus/sandbox/draft2/errorplus"

var _ errorplus.ErrorPlus = (*GoErrorPlus)(nil)

type GoErrorPlus struct {
	err error
}

// Error implements errorplus.ErrorPlus.
func (g *GoErrorPlus) Error() string {
	return g.err.Error()  // ir error is nil, this will return a panic invalid memory address or nil pointer dereference ? 
}

// Unwrap implements errorplus.ErrorPlus.
func (g *GoErrorPlus) Unwrap() error {
	panic("unimplemented")
}

// Wrap implements errorplus.ErrorPlus.
func (g *GoErrorPlus) Wrap(err error) errorplus.ErrorPlus {
	panic("unimplemented")

}
