// Go has various value types including strings,
// integers, floats, booleans, etc. Here are a few
// basic examples.

package main

import "fmt"

func main() {

	// Strings, which can be added together with `+`.
	fmt.Println("go" + "lang") // golang

	// Integers and floats.
	fmt.Println("1+1 =", 1+1)         // 1+1 = 2
	fmt.Println("7.0/3.0 =", 7.0/3.0) // 7.0/3.0 = 2.3333333333333335

	// Booleans, with boolean operators as you'd expect.
	fmt.Println(true && false) // false
	fmt.Println(true || false) // true
	fmt.Println(!true)         // false
	fmt.Println(!false)        //false
}
