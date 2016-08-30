/*
    author: diogoxiang
    time: 2016-08-16/8/15 - 23:14
    name: Learn_go
 */
package main

import "fmt"

func main() {
	var f = Adder()
	fmt.Print(f(1), " - ")
	fmt.Print(f(20), " - ")
	fmt.Print(f(300))
	/* out:
		1 - 21 - 321
		三次调用函数 f 的过程中函数 Adder() 中变量 delta 的值分别为：1、20 和 300。

		我们可以看到，在多次调用中，变量 x 的值是被保留的，即 0 + 1 = 1，然后 1 + 20 = 21，最后 21 + 300 = 321：
		闭包函数保存并积累其中的变量的值，不管外部函数退出与否，它都能够继续操作外部函数中的局部变量。

	 */
}

func Adder() func(int) int {
	var x int
	return func(delta int) int {
		x += delta
		return x
	}
}