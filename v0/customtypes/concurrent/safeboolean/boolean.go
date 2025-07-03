package safebolean

import "sync"

type GoBoolean struct {
	value       bool
	mu          sync.Mutex
}
