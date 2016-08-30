/*
Project:Learn_go
Time:2016/8/19-15:11
author:Dio
*/
package main

import (
	"fmt"
	"reflect"
)

type T struct {
	A int
	B string
	C int
	/**
	当 首字母为小写. 编译错误,  也无法 进行set 设置
	*/
}

func main() {
	t := T{23, "skidoo", 221}
	s := reflect.ValueOf(&t).Elem()
	typeOfT := s.Type()
	fmt.Println(typeOfT) //main.T
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		fmt.Printf("%d: %s %s = %v\n", i,
			typeOfT.Field(i).Name, f.Type(),f.Interface())
	}
	s.Field(0).SetInt(77)
	s.Field(1).SetString("Sunset Strip")
	s.Field(2).SetInt(233)
	fmt.Println("t is now", t)
}
