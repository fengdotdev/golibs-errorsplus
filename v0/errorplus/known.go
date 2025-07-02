package errorplus

type Known int

const (
	UnknownError Known = iota
	ExpectedError
)
