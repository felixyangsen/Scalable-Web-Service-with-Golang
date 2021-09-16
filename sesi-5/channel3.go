// package main

// import "fmt"

// func main() {
// 	message := make(chan string)
// 	valueCek := make(chan bool)

// 	go func() {
// 		message <- "ping"
// 		valueCek <- false
// 	}()

// 	go cekString(message, valueCek)

// 	resultCek := <-valueCek
// 	resultMessage := <-message

// 	fmt.Println(resultMessage, resultCek)
// }

// func cekString(message chan string, valueCek chan bool) {
// 	valueCek <- true
// 	message <- "success"
// }
