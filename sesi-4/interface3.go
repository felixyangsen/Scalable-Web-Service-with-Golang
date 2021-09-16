package main

import "fmt"

func main() {
	var temp interface{}

	temp = 1
	fmt.Println(temp)
	temp = "string"
	fmt.Println(temp)
}

