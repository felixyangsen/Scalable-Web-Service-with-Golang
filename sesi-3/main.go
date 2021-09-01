package main

import (
	"fmt"
)

func main() {
	var param1, param2 int
	var param3 string
	var param4 = []int{1, 2, 3, 4, 5, 6}

	function1()
	function2(param1, param2, param3)
	res := function3(param1, param2)
	count, err := function4(param1)
	function5(param4...)

	fmt.Println(res, err, count)
}

// tanpa param, tanpa return value
func function1() {
	fmt.Println("ini function")
}

// dengan multiple param
func function2(param1, param2 int, param3 string) {
	fmt.Println(param1, param2, param3)
}

// dengan param dan return value
func function3(param1, param2 int) int {
	return param1 + param2
}

// dengan predefined return value
func function4(param1 int) (count int, err error) {
	count = param1 + 100
	err = nil
	return
}

// dengan parameter variadic
func function5(param4 ...int) {
	fmt.Println(param4)
}

func closure1() {
	//closure sebagai variable
	var closure = func(param1 int) int {
		return param1
	}

	closure(5)
}

func closure2() {
	// immediately-invoked function expression
	var closure = func(param1 int) int {
		return param1
	}(5)

	fmt.Println(closure)
}

// closure as return value
func closure3() func() int {
	return func() int {
		return 100
	}
}

// closure as param
func closure4(tambah func(param1, param2 int) int) {
	hasil := tambah(1, 2)
	fmt.Println(hasil)
}

// closure as type
type tambah func(param1, param2 int) int

// pointer
func pointer1() {
	var a int = 4
	var b *int = &a

	fmt.Println(a)  // value a
	fmt.Println(&a) // address a
	fmt.Println(b)  // address a
	fmt.Println(&b) // address b
	fmt.Println(*b) // value b
}

// declare struct
type Human struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// embedded struct
type HighHuman struct {
	Power string
	Human
}

// nested struct
type Dungeon struct {
	Name     string
	Property struct {
		Trap   string
		Reward string
	}
}

func structure1() {
	// penggunaan struct
	var polisi Human
	// inisialisasi
	polisi.Name = "andi"
	polisi.Age = 32

	// deklarasi dan inisialisasi
	var petani = Human{"budi", 28}

	// deklarasi anonymous struct
	var highHuman = struct {
		Power string
		Human
	}{}

	// deklarasi dan inisialisasi
	// anonymous struct
	var anotherHighHuman = struct {
		Power string
		Human
	}{
		Power: "fire",
		Human: Human{
			Name: "inferno",
			Age:  22,
		},
	}

	fmt.Println(petani, highHuman, anotherHighHuman)
}

// method
func (h Human) eat() {
	fmt.Println(h.Name + " sedang makan")
}

// method pointer
func (h *Human) sleep() {
	fmt.Println(h.Name + " sedang tidur")
}
