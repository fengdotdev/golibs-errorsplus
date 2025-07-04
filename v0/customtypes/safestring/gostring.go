package safestring

import "sync"

var _ SafeString = (*GoString)(nil)

type GoString struct {
	value        string
	defaultValue string
	mu           sync.Mutex
}


