package eplus

type Severity int

const (
	Fatal Severity = iota
	Critical
	Serious
	Warning
)
