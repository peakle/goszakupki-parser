package proxy

import "strings"

// LoadProxy - load proxies to channel
func LoadProxy(queue chan<- string, done <-chan struct{}) {
	for {
		select {
		case <-done:
			return
		case queue <- strings.Replace("", "http://", "", 0):
			// TODO add proxy
		}
	}
}
