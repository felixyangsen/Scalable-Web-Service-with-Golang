// package main

// import "fmt"

// func main() {
// 	// channel as parameter
// 	i := 10
// 	ch := make(chan int, 2)

// 	go sendData(i, ch)
// 	receiveData(ch)
// }

// func sendData(idx int, ch chan int) {
// 	for i := 0; i < idx; i++ {
// 		ch <- i
// 		fmt.Println("data ke ", i)
// 	}

// 	close(ch)
// }

// func receiveData(ch chan int) {
// 	for result := range ch {
// 		fmt.Println(result)
// 	}
// }
