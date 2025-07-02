package safebolean

func (sb *GoBoolean) Stream() chan bool {

	boolChan := make(chan bool)
	sb.subscribe(boolChan)

	return boolChan
}

func (sb *GoBoolean) Await() bool {
	boolChan := sb.Stream()
	defer close(boolChan)

	select {
	case value := <-boolChan:
		return value
	}
}

func NotifyAll(value bool, subs []chan bool) {

	for _, subscriber := range subs {
		go func(value bool) {
			subscriber <- value
		}(value)
	}
}
