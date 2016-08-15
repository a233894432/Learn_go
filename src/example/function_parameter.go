/*
    author: diogoxiang
    time: 2016-08-16/8/15 - 22:55
    name: Learn_go
 */
package main

import (
	"fmt"
)

func main() {
	callback(1, Add)
//	out : The sum of 1 and 2 is: 3
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}