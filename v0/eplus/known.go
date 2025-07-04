package eplus

type Knowledge bool

const (
	UnknownError  Knowledge = false
	ExpectedError Knowledge = true
)
