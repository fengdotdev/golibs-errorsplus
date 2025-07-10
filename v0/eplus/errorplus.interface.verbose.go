package eplus

type Verbose interface {
	Verbose() string
	Trace() []struct {
		File string
		Line int
		Func string
	}
	Msg() string
	
}
