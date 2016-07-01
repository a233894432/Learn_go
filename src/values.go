package main

import (
	"fmt"
	"math"
)

//常量用const声明
const s string = "constant"

func main() {
	//字符串可以使用 + 来连接
	//	fmt.Println("go" + "lang")

	//整型和浮点数
	//	fmt.Println("1+1 =", 1+1)
	//	fmt.Println("7.0/3.0 =", 7.0/3.0)

	//布尔类型，你可以使用布尔运算符
	//	fmt.Println(true && false)
	//	fmt.Println(true || false)
	//	fmt.Println(!true)

	//用var声明一个或多个变量
	var a string = "initial"
	fmt.Println(a)
	fmt.Println(a[0])

	//声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	fmt.Println(s)
	const n = 500000000
	//数字可以在上下文指定需要的类型，例如math.Sin需要一个float64的类型
	fmt.Println(math.Sin(n))
}
