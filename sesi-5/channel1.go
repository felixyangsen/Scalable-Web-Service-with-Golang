// package main

// import "fmt"

// // biasanya digunakan dalam goroutine
// // menghubungkan goroutine satu sama lain
// // dapat mengirimkan data antar goroutine
// func main() {
// 	var message = make(chan string)

// 	var sayHelloTo = func(who string) {
// 		var data = fmt.Sprintf("hello %s", who)
// 		message <- data
// 	}

// 	go sayHelloTo("andi")
// 	go sayHelloTo("budi")
// 	go sayHelloTo("cindy")

// 	var message1 = <-message
// 	var message2 = <-message
// 	var message3 = <-message
// 	fmt.Println(message1, message2, message3)
// }
