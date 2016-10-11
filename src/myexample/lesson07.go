/*
Project:myexample
Time:2016年10月11日14:15:16
Author:diogo
*/
/* 指针

支持指针类型 *T，指针的指针 *T，以及包含包名前缀的 .T。

默认值 nil，没有 NULL 常量。
操作符 "&" 取变量地址，"*" 透过指针访问目标对象。
不支持指针运算，不支持 "->" 运算符，直接用 "." 访问目标成员。

*/

package main

import "fmt"

/** out
  字符串的类型转换
*/
func main() {
	// type data struct{ a int }

	// var d = data{1234}
	// var p *data

	// p = &d
	// fmt.Printf("%p, %v\n", p, p.a) // 直接用指针访问目标对象成员，无须转换。
	// 0x2101ef018, 1234

	// ------
	// x := 0x12345678

	// p := unsafe.Pointer(&x) // *int -> Pointer
	// n := (*[4]byte)(p)      // Pointer -> *[4]byte

	// for i := 0; i < len(n); i++ {
	// 	fmt.Printf("%X ", n[i])
	// }

	//OUT: 78 56 34 12

	// -----
	// d := struct {
	// 	s string
	// 	x int
	// }{"abc", 100}

	// p := uintptr(unsafe.Pointer(&d)) // *struct -> Pointer -> uintptr
	// p += unsafe.Offsetof(d.x)        // uintptr + offset
	// p2 := unsafe.Pointer(p)          // uintptr -> Pointer
	// px := (*int)(p2)                 // Pointer -> *int
	// *px = 200                        // d.x = 200

	// fmt.Printf("%#v\n", d)

	// out: struct { s string; x int }{s:"abc", x:200}
	type bigint int64

	x := 1234
	var b bigint = bigint(x) // 必须显式转换，除非是常量。
	var b2 int64 = int64(b)
	fmt.Println(b2)
	// out: 1234

}
