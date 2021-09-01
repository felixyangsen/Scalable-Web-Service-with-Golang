package main

import "fmt"

func main() {

}

func Variable1() {

	var a string
	var b string = "text"
	var c = "text"
	d := "text"

	fmt.Println(a, b, c, d)
}

func Variable2() {

	var a, b, c string
	var d, e, f string = "text", "text", "text"
	var g, h, i = "text", "text", "text"
	j, k, l := "text", "text", "text"

	fmt.Println(a, b, c, d, e, f, g, h, i, j, k, l)
}

func Variable3() {

	var _ = "text"
	_ = "text"

}

func Variable4() {

	var a = new(string)
	b := new(string)

	fmt.Println(a, b)

}

func Variable5() {

	c := make(chan string)
	s := make([]string, 10)
	m := make(map[string]interface{})

	fmt.Println(c, s, m)
}

func konstant() {

	const phi = 3.14

}

func Operator() {

	var a, b bool

	fmt.Println(a && b, a || b, !a)

	fmt.Println(a == b, a != b, a < b, a > b, a <= b, a >= b)

	fmt.Println(a + b, a - b, a * b, a / b, a % b)
}

// ini adalah comment

/*
	ini adalah inline comment
*/
