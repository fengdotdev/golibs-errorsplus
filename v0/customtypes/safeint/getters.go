package safeint

import "fmt"

func (si *GoInt) String() string {
	si.mu.Lock()
	defer si.mu.Unlock()
	return fmt.Sprintf("%d", si.value)
}

// Value implements SafeInt.
func (si *GoInt) Value() int {
	si.mu.Lock()
	defer si.mu.Unlock()
	return si.value
}

// Get implements SafeInt.
func (si *GoInt) Get() int {
	si.mu.Lock()
	defer si.mu.Unlock()
	return si.value
}

func (si *GoInt) GetDefault() int {
	si.mu.Lock()
	defer si.mu.Unlock()
	return si.defaultValue
}
