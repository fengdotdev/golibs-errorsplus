package errorplus

import "time"

type ErrorPlus interface {
	error
	VerboseError() error
	OnionError
}

type ErrorBase interface {
	Error() string // nil must be a panic like in the standard library
}

type SeverityError interface {
	Severity() string
	HasSeverity() bool
}

type OnionError interface {
	Surface() error
	Origin() error
	Wrap(err error) error
	Unwrap() error
	Deep() int
	DeepLevel(int)
}

type OnOnionError interface {
	IsOrigin() bool
	IsSurface() bool
	HasChild() bool
}

type ErrorPlusExtension interface {
	Trace() []string
	TimeStamp() time.Time
	Source() string
	SourceLine() int
	SourceFunc() string
	SourceFile() string
	SourcePackage() string
	Tags() []string
	Tag(string) string
	HasTag(string) bool
	HasTags() bool
	Message() string
	HasMessage() bool
	Fn() string
	HasFn() bool
	FnArgs() []interface{}
	HasFnArgs() bool
	ObscureArgs() bool
}
