/*
Project:myexample
Time:2016年10月11日17:05:00
Author:diogo
*/

/*

3.1 函数定义
不支持 嵌套 (nested)、重载 (overload) 和 默认参数 (default parameter)。

无需声明原型。
支持不定长变参。
支持多返回值。
支持命名返回参数。
支持匿名函数和闭包。

*/

package main

import "fmt"

func test(fn func() int) int {
	return fn()
}

type FormatFunc func(s string, x, y int) string // 定义函数类型。

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func main() {
	s1 := test(func() int { return 100 }) // 直接将匿名函数当参数。

	s2 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	println(s1, s2)
}
