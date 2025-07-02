package errorplus

type Severity int

const (
	Fatal Severity = iota
	Critical
	Serious
	Warning
)
