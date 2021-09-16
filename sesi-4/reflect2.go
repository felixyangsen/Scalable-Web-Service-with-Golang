package main

import (
	"fmt"
	"reflect"
)

func main() {
	var andi = &student{Name: "andi", Grade: 2}
	andi.getPropertyInfo()
}

type student struct {
	Name  string
	Grade int
}

func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println(reflectType.Field(i).Name)
		fmt.Println(reflectType.Field(i).Type)
		fmt.Println(reflectValue.Field(i).Interface())
	}
}

