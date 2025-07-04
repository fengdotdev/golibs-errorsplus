package safebolean

import "sync"

func True() *GoBoolean {
	return &GoBoolean{
		value:        true,
		defaultValue: true,
		mu:           sync.Mutex{},
	}
}

func False() *GoBoolean {
	return &GoBoolean{
		value:        false,
		defaultValue: false,
		mu:           sync.Mutex{},
	}
}

func New(value bool, defaultValue bool) *GoBoolean {
	return &GoBoolean{
		value:        value,
		defaultValue: defaultValue,
		mu:           sync.Mutex{},
	}
}
