package eplus

import "time"

type ErrorPlusExtension interface {
	ErrorPlusTrace
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

type ErrorPlusTrace interface {
	// Trace returns a slice of strings representing the trace of the error. the slice is striped of runtime-callers and the package errorplus.
	Trace() []string
	// TraceRaw returns a slice of Trace objects representing the raw trace of the error without any modifications.
	TraceRaw() []Trace
	// TraceString returns a string representation of the trace of the error.
	// The string is a concatenation of the trace strings. striped of runtime-callers and the package errorplus.
	TraceString() string
	// TraceFull returns a slice of strings representing the full trace of the error. with runtime-callers and the package errorplus.
	TraceFull() []string
}
