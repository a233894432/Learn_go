/*
   author: diogoxiang
   time: 2016-08-16/8/26 - 23:56
   name: Learn_go
*/

/**
代码书写
defer -> recover => panic

发生错误后恢复>程序
panic -> defer -> recover
 */
package main

import (
	"fmt"
)

func badCall() {
	panic("bad end")
}

func test() {
	defer func() {
		if e := recover(); e != nil {
			fmt.Printf("Panicing %s\r\n", e)
		}
	}()
	badCall()
	fmt.Printf("After bad call\r\n") // <-- wordt niet bereikt
}

func main() {
	fmt.Printf("Calling test\r\n")
	test()
	fmt.Printf("Test completed\r\n")
}
