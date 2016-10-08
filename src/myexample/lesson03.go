/*
Project:myexample
Time:2016/9/30-14:20
Author:diogo
*/
package main

func main() {
	//支持在函数内 goto 跳转。标签名区分大小写，未使用标签引发错误。
	var i int
	for {
		println(i)
		i++
		if i > 2 {
			goto BREAK
		}
	}
BREAK:
	println("break")

//EXIT: // Error: label EXIT defined and not used
}
