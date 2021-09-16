package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(3)
	go foo(3, &wg)
	go foo(2, &wg)
	go foo(1, &wg)
	wg.Wait()
}

func foo(d time.Duration, wg *sync.WaitGroup) {
	d *= 10000000000
	time.Sleep(d)
	fmt.Println(d)
	wg.Done()
}
