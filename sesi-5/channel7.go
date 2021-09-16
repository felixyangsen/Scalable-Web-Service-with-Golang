// package main

// import "fmt"

// func main() {
// 	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
// 	ch := make(chan int, 3)

// 	go collect(data, ch)
// 	cetak(ch)
// }

// func collect(data []int, ch chan int) {
// 	for _, dt := range data {
// 		ch <- dt
// 	}

// 	close(ch)
// }

// func cetak(ch chan int) {
// 	for result := range ch {
// 		fmt.Println(result)
// 	}
// }
