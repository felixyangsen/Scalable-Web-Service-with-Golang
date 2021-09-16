package main

import (
	"fmt"
	"reflect"
)

func main() {
	var andi = &student{Name: "andi", Grade: 2}

	var reflectValue = reflect.ValueOf(andi)
	var method = reflectValue.MethodByName("setName")
	method.Call([]reflect.Value{
		reflect.ValueOf("budi"),
	})

	fmt.Println(andi.Name)
}

type student struct {
	Name  string
	Grade int
}

func (s *student) setName(name string) {
	s.Name = name
}

