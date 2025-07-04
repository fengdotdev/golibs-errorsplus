package safeint

import "sync"



// Return a new instance of GoInt with the given value and default value.
// The default value is used when Reset is called.
// The mutex is initialized to ensure thread safety.
func New(value int) *GoInt {
	return &GoInt{
		value:        value,
		defaultValue: value,
		mu:           sync.Mutex{},
	}
}

func NewWithDefault(value, defaultValue int) *GoInt {
	return &GoInt{
		value:        value,
		defaultValue: defaultValue,
		mu:           sync.Mutex{},
	}
}
