/*
   author: diogoxiang
   time: 2016-08-16/8/15 - 23:48
   name: Learn_go
*/

/**
6.12 通过内存缓存来提升性能
内存缓存的技术在使用计算成本相对昂贵的函数时非常有用（不仅限于例子中的递归），
譬如大量进行相同参数的运算。这种技术还可以应用于纯函数中，即相同输入必定获得相同输出的函数。



要计算数列中第 n 个数字，需要先得到之前两个数的值，但很明显绝大多数情况下前两个数的值都是已经计算过的。
即每个更后面的数都是基于之前计算结果的重复计算，正如示例 6.11 fibonnaci.go 所展示的那样。

而我们要做就是将第 n 个数的值存在数组中索引为 n 的位置（详见第 7 章），
然后在数组中查找是否已经计算过，如果没有找到，则再进行计算。

*/
package main

import (
	"fmt"
	"time"
)

const LIM = 41

var fibs [LIM]uint64

func main() {
	var result uint64 = 0
	start := time.Now()
	for i := 0; i < LIM; i++ {
		result = fibonacci(i)
		fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	}
	end := time.Now()
	delta := end.Sub(start)
	fmt.Printf("longCalculation took this amount of time: %s\n", delta)
}
func fibonacci(n int) (res uint64) {
	// memoization: check if fibonacci(n) is already known in array:
	if fibs[n] != 0 {
		res = fibs[n]
		return
	}
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	fibs[n] = res
	return
}
