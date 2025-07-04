package safeint

// Decrement implements SafeInt.
func (si *GoInt) Decrement() {
	si.mu.Lock()
	defer si.mu.Unlock()
	si.value--
}

// Increment implements SafeInt.
func (si *GoInt) Increment() {
	si.mu.Lock()
	defer si.mu.Unlock()
	si.value++
}

// Reset implements SafeInt.
func (si *GoInt) Reset() {
	si.mu.Lock()
	defer si.mu.Unlock()
	si.value = si.defaultValue
}

// Set implements SafeInt.
func (si *GoInt) Set(value int) {
	si.mu.Lock()
	defer si.mu.Unlock()
	si.value = value
}
