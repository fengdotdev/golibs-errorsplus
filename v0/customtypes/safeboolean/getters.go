package safebolean

func (sb *GoBoolean) String() string {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	if sb.value {
		return "true"
	}
	return "false"
}

func (sb *GoBoolean) Get() bool {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	return sb.value
}

func (sb *GoBoolean) Value() bool {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	return sb.value
}

func (sb *GoBoolean) GetDefault() bool {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	return sb.defaultValue
}
