/*
Project:Learn_go
Time:2016/8/17-14:00
author:Dio
*/
/*
使用 new

使用 new 函数给一个新的结构体变量分配内存，它返回指向已分配内存的指针：var t *T = new(T)，
如果需要可以把这条语句放在不同的行（比如定义是包范围的，但是分配却没有必要在开始就做）。

var t *T
t = new(T) 写这条语句的惯用方法是：t := new(T)，变量 t 是一个指向 T的指针，此时结构体字段的值是它们所属类型的零值。

声明 var t T 也会给 t 分配内存，并零值化内存，但是这个时候 t 是类型T。
在这两种方式中，t 通常被称做类型 T 的一个实例（instance）或对象（object）。
 */
package main
import "fmt"

type struct1 struct {
	i1  int
	f1  float32
	str string
}

func main() {
	ms := new(struct1)
	ms.i1 = 10
	ms.f1 = 15.5
	ms.str= "Chris"

	fmt.Printf("The int is: %d\n", ms.i1)
	fmt.Printf("The float is: %f\n", ms.f1)
	fmt.Printf("The string is: %s\n", ms.str)
	fmt.Println(ms)
}

/** OUT:
The int is: 10
The float is: 15.500000
The string is: Chris
&{10 15.5 Chris}
 */
