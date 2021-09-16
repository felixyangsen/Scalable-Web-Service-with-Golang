// package main

// import "fmt"

// func main() {
// 	// slice channel
// 	c := make(chan []int)
// 	var things []int
// 	i := 0

// 	for {
// 		things = append(things, i)
// 		if i == 9 {
// 			break
// 		}
// 		i++
// 	}

// 	go func() {
// 		c <- things
// 	}()

// 	for _, j := range <-c {
// 		fmt.Println(j)
// 	}
// }
