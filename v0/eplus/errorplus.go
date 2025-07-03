package eplus

type ErrorPlus interface {
	error
	Verbose
	//OnionError
}
