/*
Project:Learn_go
Time:2016/8/18-14:28
author:Dio
*/
/*
当定义了一个有很多方法的类型时，十之八九你会使用 String() 方法来定制类型的字符串形式的输出，
换句话说：一种可阅读性和打印性的输出。如果类型定义了 String() 方法，
它会被用在 fmt.Printf() 中生成默认的输出：等同于使用格式化描述符 %v 产生的输出。
还有 fmt.Print() 和 fmt.Println() 也会自动使用 String() 方法。
*/
package main

import (
	"fmt"
	"strconv"
)

type TwoInts struct {
	a int
	b int
}
type T struct {
	a int
	b float32
	c string
}

func (t *T) String() string {

	return strconv.Itoa(t.a) + "/" +strconv.FormatFloat(float64(t.b), 'f', 3, 64) + "/" + t.c

}

func main() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10
	fmt.Printf("two1 is: %v\n", two1)
	fmt.Println("two1 is:", two1)
	fmt.Printf("two1 is: %T\n", two1)
	fmt.Printf("two1 is: %#v\n", two1)

	fmt.Println()
	fmt.Printf("two1.a is: %v\n", two1.a)
	fmt.Println()
	t := &T{7, 3.23, "diogo"}
	fmt.Printf("%v\n", t)

	fmt.Println()
	f, _ := strconv.ParseFloat("3.1415", 64)
	fmt.Println(f)

}

func (tn *TwoInts) String() string {
	return "(" + strconv.Itoa(tn.a) + "/" + strconv.Itoa(tn.b) + ")"
}
