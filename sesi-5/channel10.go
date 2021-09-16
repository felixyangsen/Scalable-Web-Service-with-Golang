// package main

// import "fmt"

// func main() {
// 	result1 := make(chan string)
// 	result2 := make(chan string)

// 	go func() {
// 		print1(result1)
// 	}()

// 	go func() {
// 		print2(result2)
// 	}()

// 	select {
// 	case c1 := <-result1:
// 		fmt.Println(c1)
// 	case c2 := <-result2:
// 		fmt.Println(c2)
// 	}
// }

// func print1(ch chan string) {
// 	ch <- "format 1"
// }

// func print2(ch chan string) {
// 	ch <- "format 2"
// }
