package main

import (
	"fmt"
	"reflect"
)

func main() {
	value := 1
	reflectValue := reflect.ValueOf(value)

	// tipe data dari variable
	fmt.Println(reflectValue.Type())

	// nilai dari variable
	fmt.Println(reflectValue.Int())
}

