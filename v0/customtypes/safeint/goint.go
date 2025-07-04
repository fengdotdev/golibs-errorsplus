package safeint

import "sync"

var _ SafeInt = (*GoInt)(nil)

type GoInt struct {
	value        int
	defaultValue int
	mu           sync.Mutex
}
