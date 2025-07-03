package errorplus

type SeverityError interface {
	Severity() string
	HasSeverity() bool
}
