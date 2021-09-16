// package main

// import (
// 	"fmt"
// 	"time"
// )

// func main() {
// 	ch := make(chan int, 3)
// 	go write(ch)
// 	time.Sleep(3 * time.Second)
// 	for v := range ch {
// 		fmt.Println("read ", v)
// 		time.Sleep(2 * time.Second)
// 	}
// }

// func write(ch chan int) {
// 	for i := 0; i < 10; i++ {
// 		ch <- i
// 		fmt.Println("write ", i)
// 	}

// 	close(ch)
// }
