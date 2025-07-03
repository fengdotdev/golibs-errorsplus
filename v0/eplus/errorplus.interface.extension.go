package eplus

import "time"

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
