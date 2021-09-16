// // package main

// import "fmt"

// func main() {
// 	c := make(chan string, 1)
// 	c <- "receive only"

// 	f(c)
// }

// func f(c <-chan string) {
// 	message := <-c
// 	fmt.Println(message)
// }
