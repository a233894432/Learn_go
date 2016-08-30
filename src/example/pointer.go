//这个是用于学习 指针
package main

import "fmt"

func main() {

	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)

	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	/**
	指针类型的一些说明
	*/
	var a int = 1
	var b *int = &a
	var c **int = &b
	var x int = *b
	fmt.Println("a = ", a)
	fmt.Println("&a = ", &a)
	fmt.Println("*&a = ", *&a)
	fmt.Println("b = ", b)
	fmt.Println("&b = ", &b)
	fmt.Println("*&b = ", *&b)
	fmt.Println("*b = ", *b)
	fmt.Println("c = ", c)
	fmt.Println("*c = ", *c)
	fmt.Println("&c = ", &c)
	fmt.Println("*&c = ", *&c)
	fmt.Println("**c = ", **c)
	fmt.Println("***&*&*&*&c = ", ***&*&*&*&*&c)
	fmt.Println("x = ", x)
	/*
	a     =     1
	&a     =     0xc200000018
	*&a     =     1
	b     =     0xc200000018
	&b     =     0xc200000020
	*&b     =     0xc200000018
	*b     =     1
	c     =     0xc200000020
	*c     =     0xc200000018
	&c     =     0xc200000028
	*&c     =     0xc200000020
	**c     =     1
	***&*&*&*&c     =     1
	x     =     1
	*/

}
