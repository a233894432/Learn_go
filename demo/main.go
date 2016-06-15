package main

import (
	"fmt"
)

const c = "C"

var v int = 5

type T struct{}

func init() { // initialization of package
	fmt.Println("sss")
}

func main() {
	var a int
	Func1()
	// ...
	fmt.Println(a)
}

func (t T) Method1() {
	//...
}

func Func1() { // exported function Func1
	//...
	fmt.Println("Fun1")
	fmt.Println(".....")
	fmt.Println("...")
	var arr [10]int                                 // 声明了一个int类型的数组
	arr[0] = 42                                     // 数组下标是从0开始的
	arr[1] = 13                                     // 赋值操作
	fmt.Printf("The first element is %d\n", arr[0]) // 获取数据，返回42
	fmt.Printf("The last element is %d\n", arr[9])  //返回未赋值的最后一个元素，默认返回0

	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	// 声明两个含有byte的slice
	var a, b []byte

	// a指向数组的第3个元素开始，并到第五个元素结束，
	a = ar[2:5]
	//现在a含有的元素: ar[2]、ar[3]和ar[4]

	fmt.Printf("The a  element is %d\n", a)
	// b是数组ar的另一个slice

	b = ar[3:5]
	// b的元素是：ar[3]和ar[4]
	fmt.Printf("The b  element is %d\n", b)

}
