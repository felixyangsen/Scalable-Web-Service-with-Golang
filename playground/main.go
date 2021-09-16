package main

import "fmt"

func main() {

}

func ifEx() {
	// if
	var y = 16
	if y >= 17 {
		fmt.Println("y bigger than 17")
	} else if y <= 6 {
		fmt.Println("y smaller than 6")
	} else {
		fmt.Println("we don't know")
	}
}

func forEx() {
	// for loop
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	for {
		fmt.Println("infinite")
	}
}

func switchEx() {
	// switch case
	x := 2
	switch x {
	case 1:
		fmt.Println("x is 1")
	case 2:
		fmt.Println("x is 2")
	default:
		fmt.Println("don't know")
	}
}

func scanEx() {
	var input int
	fmt.Printf("input: ")
	fmt.Scan(&input)
	fmt.Println("input = ", input)
}
