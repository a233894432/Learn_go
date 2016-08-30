/*
   author: diogoxiang
   time: 2016-08-16/8/15 - 23:10
   name: Learn_go
*/
/*
6.9 应用闭包：将函数作为返回值
*/

package main

import "fmt"

func main() {
	// make an Add2 function, give it a name p2, and call it:
	p2 := Add2()
	fmt.Printf("Call Add2 for 3 gives: %v\n", p2(3))
	// make a special Adder function, a gets value 3:
	TwoAdder := Adder(2)
	fmt.Printf("The result is: %v\n", TwoAdder(3))

	/*
	Call Add2 for 3 gives: 5
	The result is: 5
	*/

}

func Add2() func(b int) int {
	return func(b int) int {
		return b + 2
	}
}

func Adder(a int) func(b int) int {
	return func(b int) int {
		return a + b
	}
}
