package safebolean

import "sync"

var _ SafeBoolean = (*GoBoolean)(nil)

type GoBoolean struct {
	value        bool
	defaultValue bool
	mu           sync.Mutex
}
