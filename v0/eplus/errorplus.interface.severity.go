package eplus

type SeverityError interface {
	Severity() string
	HasSeverity() bool
}
