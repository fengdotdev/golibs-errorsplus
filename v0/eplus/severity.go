package eplus

type Severity int

const (
	UnknownSev  Severity = 0 // default severity
	FatalSev    Severity = 1
	CriticalSev Severity = 2
	SeriousSev  Severity = 3
	WarningSev  Severity = 4
)
