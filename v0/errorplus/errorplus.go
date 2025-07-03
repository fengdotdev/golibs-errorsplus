package errorplus

type ErrorPlus interface {
	error
	Verbose
	//OnionError
}
