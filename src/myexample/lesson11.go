/*
Project:myexample
Time:2016年10月12日10:38:54
Author:diogo
*/
package main

import (
	"errors"
	"fmt"
)

var ErrDivByZero = errors.New("division by zero")

func div(x, y int) (int, error) {
	if y == 0 {
		return 0, ErrDivByZero
	}
	return x / y, nil
}

func main() {
	switch z, err := div(10, 21); err {
	case nil:
		println(z)
	case ErrDivByZero:
		panic(err)
	}

	a := [2]int{}
	fmt.Printf("a: %p\n", &a)
	test(a)
	fmt.Println(a)

	/* out
	a: 0x2101f9150
	x: 0x2101f9170
	[0 0]

	*/

}

func test(x [2]int) {
	fmt.Printf("x: %p\n", &x)
	x[1] = 1000
}
