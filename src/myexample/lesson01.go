/*
Project:myexample
Time:2016/9/30-11:19
Author:diogo
*/
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//s := "abc"
	//fmt.Println(s[0] == '\x61', s[1] == 'b', s[2] == 0x63)

	//s := "abcd"
	//bs := []byte(s)
	//
	//bs[1] = 'B'
	//fmt.Println(string(bs))
	//
	//u := "电脑"
	//us := []rune(u)
	//
	//us[1] = '话'
	//println(string(us))
	//type data struct{ a int }
	//
	//var d = data{1234}
	//var p *data
	//
	//p = &d
	//fmt.Printf("%p, %v\n", p, p.a) // 直接用指针访问目标对象成员，无须转换。

	d := struct {
		s string
		x int
	}{"abc", 100}

	p := uintptr(unsafe.Pointer(&d)) // *struct -> Pointer -> uintptr
	p += unsafe.Offsetof(d.x) // uintptr + offset
	p2 := unsafe.Pointer(p) // uintptr -> Pointer
	px := (*int)(p2) // Pointer -> *int
	*px = 200 // d.x = 200

	fmt.Printf("%#v\n", d)  // struct { s string; x int }{s:"abc", x:200}
}
