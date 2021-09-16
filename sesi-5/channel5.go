// package main

// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// func main() {
// 	runtime.GOMAXPROCS(2)
// 	sliceChan := make(chan []int)

// 	arrInt := []int{}

// 	var wg sync.WaitGroup
// 	var mtx sync.Mutex

// 	x := 0

// 	for i := 0; i < 1000; i++ {
// 		wg.Add(1)
// 		go func() {
// 			mtx.Lock()

// 			x++
// 			arrInt = append(arrInt, x)

// 			mtx.Unlock()

// 			defer wg.Done()
// 		}()
// 	}

// 	wg.Wait()

// 	go func() {
// 		sliceChan <- arrInt
// 	}()

// 	for _, rangeA := range <-sliceChan {
// 		fmt.Println(rangeA)
// 	}
// }
