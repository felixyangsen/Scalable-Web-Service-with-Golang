package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string = "ok"
	defer catch()

	_, err := strconv.Atoi(str)

	if err != nil {
		panic(err)
	}
}

func catch() {
	r := recover()

	if r != nil {
		fmt.Println("error occured", r)
	} else {
		fmt.Println("done")
	}
}
