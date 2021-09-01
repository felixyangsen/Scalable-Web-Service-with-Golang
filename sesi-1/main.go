package main

import (
	"fmt"
)

func main() {
	learnVariable()
	learnVariable2()
	multiVariable()
	konstanta()
	quiz()
}

func learnVariable() {
	var str string
	str = "learn golang"

	var strr = "learn golang"

	fmt.Println(str)
	fmt.Println(strr)

	fmt.Printf("sedang belajar %s", "golang")
}

func learnVariable2() {
	s := "string"
	d := 1

	fmt.Println(s, d)
}

func multiVariable() {
	// this is a comment
	var a, b, c string

	a = "a"
	b = "b"
	c = "c"

	fmt.Println(a, b, c)
}

func konstanta() {
	const phi = 3.14

	fmt.Println(phi)
}

func quiz() {
	var a string
	a = "testing1"

	fmt.Println(a)
}
