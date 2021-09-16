// package main

// import "fmt"

// func main() {
// 	canSleep := make(chan bool)

// 	go mauTidur("pacar", canSleep)
	
// 	if <-canSleep {
// 		fmt.Println("bisa tidur")
// 	} else {
// 		fmt.Println("gabisa tidur")
// 	}
// }

// func mauTidur(bebanPikiran string, canSleep chan bool) {
// 	if bebanPikiran == "tugas" {
// 		canSleep <- false
// 	} else if bebanPikiran == "pacar" {
// 		canSleep <- false
// 	} else {
// 		canSleep <- true
// 	}
// }
