package main

import (
	"fmt"
	"reflect"
)

type ta struct {
	name string
	age  int
}

func main() {
	bobo := ta{"diogobobo", 20}

	fmt.Println(bobo)

	t := reflect.ValueOf(&bobo)
	fmt.Println(t.Type())
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("value:", v.Float())

}
