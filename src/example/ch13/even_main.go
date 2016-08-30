/*
    author: diogoxiang
    time: 2016-08-16/8/30 - 01:11
    name: Learn_go
 */
package main

import (
	"fmt"
	"example/ch13/even"
)

func main() {

	for i:=0; i<=100; i++ {
		fmt.Printf("Is the integer %d even? %v\n", i, even.Even(i))
	}
}
