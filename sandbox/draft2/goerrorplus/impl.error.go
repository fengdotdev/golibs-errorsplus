package goerrorplus

var _ error = (*GoErrorPlus)(nil)

func (g *GoErrorPlus) Error() string {
	return g.err.Error() // ir error is nil, this will return a panic invalid memory address or nil pointer dereference ?
}
