/*
Project:Learn_go
Time:2016/8/19-14:58
author:Dio
*/
package main

import (
	"fmt"
	"reflect"
)

type NotknownType struct {
	s1, s2, s3 string
}

type Dioname struct {
	d1 int
}


func (n NotknownType) String() string {
	return n.s1 + " - " + n.s2 + " - " + n.s3
}

// variable to investigate:
var secret interface{} = NotknownType{"Ada", "Go", "eberon"}

func main() {
	value := reflect.ValueOf(secret) // <main.NotknownType Value>
	typ := reflect.TypeOf(secret)    // main.NotknownType
	// alternative:
	//typ := value.Type()  // main.NotknownType
	fmt.Println(value) //{Ada Go Oberon}
	fmt.Println(typ) // main.NotknownType
	knd := value.Kind() // struct
	fmt.Println(knd) // 类型

	vd:=reflect.ValueOf(Dioname{1})
	fmt.Println(vd)// {1}
	fmt.Println(vd.Kind())

	// iterate through the fields of the struct:
	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field %d: %v\n", i, value.Field(i))
		// error: panic: reflect.Value.SetString using value obtained using unexported field
		//value.Field(i).SetString("C#")
	}

	// call the first method, which is String():
	results := value.Method(0).Call(nil)
	fmt.Println(results) // [Ada - Go - Oberon]
}
