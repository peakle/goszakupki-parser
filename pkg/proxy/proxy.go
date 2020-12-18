package proxy

// LoadProxy - load proxies to channel
func LoadProxy(queue chan<- string, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		case queue <- "":
			// TODO add proxy
		}
	}
}
