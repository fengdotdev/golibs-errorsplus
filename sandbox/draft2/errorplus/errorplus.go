package errorplus

type ErrorPlus interface {
	Error() string // nil must be a panic like in the standard library
	Unwrap() error
	Wrap(err error) ErrorPlus
}
