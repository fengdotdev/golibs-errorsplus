package safestring

// Set implements SafeString.
func (g *GoString) Set(value string) {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.value = value
}

// Reset implements SafeString.
func (g *GoString) Reset() {
	g.mu.Lock()
	defer g.mu.Unlock()
	g.value = g.defaultValue
}
