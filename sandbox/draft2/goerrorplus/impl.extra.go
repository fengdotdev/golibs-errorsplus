package goerrorplus

// Unwrap implements errorplus.ErrorPlus.
func (g *GoErrorPlus) Unwrap() error {
	panic("unimplemented")
}

// Wrap implements errorplus.ErrorPlus.
func (g *GoErrorPlus) Wrap(err error) error {
	panic("unimplemented")

}
