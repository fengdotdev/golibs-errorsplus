package errorplus

type GoError struct {
	child     error
	selfError error
	msg       string
	trace     []string
}

