/*
Project:Learn_go
Time:2016/8/17-16:49
author:Dio
*/
package main

import "fmt"

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

func main() {
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
	/* out
	1 2 3 4
	{1 2}
	 */
}
