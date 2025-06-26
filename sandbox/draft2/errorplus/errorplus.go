package errorplus

type ErrorPlus interface {
	Error() string
	Unwrap() error
	Wrap(err error) ErrorPlus
}
