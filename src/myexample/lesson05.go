/*
Project:myexample
Time:2016/10/10-14:19
Author:diogo
*/
package main

import "fmt"

/** out
字符串的类型转换
*/
func main() {
	fmt.Printf("%T\n", 'a') // int32  rune 是 int32 的别名

	var c1, c2 rune = '\u6211', '们'
	println(c1 == '我', string(c2) == "\xe4\xbb\xac") // true true

	s := "abcd"
	bs := []byte(s)

	bs[1] = 'B'
	println(string(bs))

	u := "电脑"
	us := []rune(u)

	us[1] = '话'
	println(string(us))
}
