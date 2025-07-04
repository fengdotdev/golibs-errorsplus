package safestring

// String implements SafeString.
func (g *GoString) String() string {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.value
}

// Get implements SafeString.
func (g *GoString) Get() string {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.value
}

// GetDefault implements SafeString.
func (g *GoString) GetDefault() string {
	g.mu.Lock()
	defer g.mu.Unlock()
	return g.defaultValue
}
