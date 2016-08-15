/*
   author: diogoxiang
   time: 2016-08-16/8/15 - 23:30
   name: Learn_go
*/
/*
6.10 使用闭包调试
*/
package main

import (
	"fmt"
	"log"
)

func main() {

	func1()
}

var where = log.Print

func func1() {
	where()
	fmt.Println("------1")
	where()
	fmt.Println("------2")
	where()
}
