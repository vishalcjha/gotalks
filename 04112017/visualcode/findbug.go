package main

import (
	"sync"
)

func main() {
	m := sync.Mutex{}
	m.Lock()
	for i := 0; i < 10; i++ {
		defer m.Unlock()
	}
}
