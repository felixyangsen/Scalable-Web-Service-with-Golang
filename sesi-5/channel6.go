// package main

// import (
// 	"fmt"
// 	"runtime"
// )

// func main() {
// 	runtime.GOMAXPROCS(2)

// 	var message = make(chan string)

// 	listData := []string{"a", "b", "c"}

// 	for _, each := range listData {
// 		go func(who string) {
// 			var data = fmt.Sprintf("hello %s", who)
// 			message <- data
// 		}(each)
// 	}

// 	for i := 0; i < len(listData); i++ {
// 		PrintMessage(message)
// 	}
// }

// func PrintMessage(what chan string) {
// 	fmt.Println(<-what)
// }
