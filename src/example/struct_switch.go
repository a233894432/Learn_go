/*
Project:Learn_go
Time:2016/8/17-14:38
author:Dio
*/
package main

import (
	"fmt"
)

type Element interface {}

func main() {
	var e Element = 100
	switch value := e.(type) {
	case int:
		fmt.Println("int", value)
	case string:
		fmt.Println("string", value)
	default:
		fmt.Println("unknown", value)
	}
}
