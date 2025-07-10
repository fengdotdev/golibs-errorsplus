package eplus

import "time"

type ErrorPlus interface {
	error
	Verbose() string
	Trace() []struct {
		File string
		Line int
		Func string
	}
	Msg() string
	TimeStamp() time.Time
	Severity() string

	IsTemporal() bool // true if the error is temporal, false if it is permanent 
	IsUnknown() bool // true if the error is unknown, false if it is expected

	Tags() []string // ordered tag
	Unwrap() error
	Wrap(error)
	Fn() string
	FnArgs() map[string]interface{}
	HasFnArgs() bool
}
