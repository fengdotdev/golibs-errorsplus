package safestring

import "sync"

func New(str string) SafeString {
	return &GoString{
		value:        str,
		defaultValue: str,
		mu:           sync.Mutex{},
	}
}

func NewWithDefault(str, defaultStr string) SafeString {
	return &GoString{
		value:        str,
		defaultValue: defaultStr,
		mu:           sync.Mutex{},
	}
}
