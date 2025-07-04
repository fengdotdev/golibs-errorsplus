package safebolean

func (sb *GoBoolean) Set(value bool) {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.value = value
}

func (sb *GoBoolean) Toggle() {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.value = !sb.value

}

func (sb *GoBoolean) Reset() {
	sb.mu.Lock()
	defer sb.mu.Unlock()
	sb.value = sb.defaultValue
}
