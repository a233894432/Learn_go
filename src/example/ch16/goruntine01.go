/*
Project:ch16
Time:2016/9/7-18:10
Author:diogo
*/
package main

import (
	"fmt"
	"time"
)

var values = [5]int{10, 11, 12, 13, 14}

func main() {
	// 版本A:
	for ix := range values { // ix是索引值
		func() {
			fmt.Print(ix, " ")
		}() // 调用闭包打印每个索引值
	}
	fmt.Println()
	// 版本B: 和A版本类似，但是通过调用闭包作为一个协程
	for ix := range values {
		go func() {
			fmt.Print(ix, " ")
		}()
	}
	fmt.Println()
	time.Sleep(5e9)
	// 版本C: 正确的处理方式
	for ix := range values {
		go func(ix interface{}) {
			fmt.Print(ix, " ")
		}(ix)
	}
	fmt.Println()
	time.Sleep(5e9)
	// 版本D: 输出值:
	for ix := range values {
		val := values[ix]
		go func() {
			fmt.Print(val, " ")
		}()
	}
	time.Sleep(1e9)
}

/* 输出：

        0 1 2 3 4

        4 4 4 4 4

        1 0 3 4 2

        10 11 12 13 14
*/

/**
版本A调用闭包5次打印每个索引值，版本B也做相同的事，但是通过协程调用每个闭包。按理说这将执行得更快，

因为闭包是并发执行的。如果我们阻塞足够多的时间，让所有协程执行完毕，版本B的输出是：4 4 4 4 4。为什么会这样？
在版本B的循环中，ix变量 实际是一个单变量，表示每个数组元素的索引值。因为这些闭包都只绑定到一个变量，
这是一个比较好的方式，当你运行这段代码时，你将看见每次循环都打印最后一个索引值4，而不是每个元素的索引值。
因为协程可能在循环结束后还没有开始执行，而此时ix值是4。

版本C的循环写法才是正确的：调用每个闭包是将ix作为参数传递给闭包。ix在每次循环时都被重新赋值，
并将每个协程的ix放置在栈中，所以当协程最终被执行时，每个索引值对协程都是可用的。
注意这里的输出可能是0 2 1 3 4或者0 3 1 2 4或者其他类似的序列，这主要取决于每个协程何时开始被执行。

在版本D中，我们输出这个数组的值，为什么版本B不能而版本D可以呢？

因为版本D中的变量声明是在循环体内部，所以在每次循环时，这些变量相互之间是不共享的，
所以这些变量可以单独的被每个闭包使用。
 */