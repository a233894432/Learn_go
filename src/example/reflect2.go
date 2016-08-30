/*
Project:Learn_go
Time:2016/8/19-14:40
author:Dio
*/
package main

import (
	"fmt"
	"reflect"
)

func main() {
	var x float64 = 3.4
	v := reflect.ValueOf(x)
	// setting a value:
	// v.SetFloat(3.1415) // Error: will panic: reflect.Value.SetFloat using unaddressable value
	fmt.Println("settability of v:", v.CanSet())  // out : settability of v: false
	v = reflect.ValueOf(&x) // Note: take the address of x.
	fmt.Println("type of v:", v.Type()) // out :  type of v: *float64
	fmt.Println("settability of v:", v.CanSet()) // out: settability of v: false
	v = v.Elem() //转换
	fmt.Println("The Elem of v is: ", v) // out: The Elem of v is:  <float64 Value>
	fmt.Println("settability of v:", v.CanSet()) // out: settability of v: true
	v.SetFloat(3.1415) // this works!
	fmt.Println(v.Interface()) // 3.1415
	fmt.Println(v) // 3.1415
}
