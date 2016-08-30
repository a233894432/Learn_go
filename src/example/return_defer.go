/*
   author: diogoxiang
   time: 2016-08-16/8/15 - 23:08
   name: Learn_go
*/
/*
6.8
请学习以下示例并思考（return_defer.go）：函数 f 返回时，变量 ret 的值是什么？
变量 ret 的值为 2，因此 ret++，这是在执行 reutrn 1 语句后发生的。
*/
package main

import "fmt"

func f() (ret int) {
	defer func() {
		ret++
	}()
	return 1
}
func main() {
	fmt.Println(f())
}
