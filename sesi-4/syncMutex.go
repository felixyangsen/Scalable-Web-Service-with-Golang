package main

import (
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(2)

	var wg sync.WaitGroup
	var meter counter
	var mtx sync.Mutex

	for i := 0; i < 1000; i++ {
		wg.Add(1)

		go func() {
			for j := 0; j < 1000; j++ {
				mtx.Lock()
				meter.add()
				mtx.Unlock()
			}

			wg.Done()
		}()
	}
}

type counter struct {
	val int
}

func (c *counter) add() {
	c.val++
}

func (c *counter) value() (x int) {
	return c.val
}
