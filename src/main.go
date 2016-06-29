package main

import (
	"fmt"
	"reflect" //用于查看数据类型

	//	"mymath"
)

func main() {
	//	fmt.Printf("Hello, wssorld.  Sqrt(2) = %v\n", mymath.Sqrt(2))
	//	test1()
	//	test2()
	//	test3()
	//	test4()
	test5()
}

func test1() {
	sum := 1
	for sum < 10 {
		sum += sum
	}
	fmt.Printf("%v\n", sum)
}

func test2() {
	for index := 10; index > 0; index-- {
		if index == 2 {
			//			break // 或者continue
			continue
		}
		fmt.Println(index)
	}

}

func test3() {
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3 or 4")
	case 10:
		fmt.Println("i is equal to 10")
	default:
		fmt.Println("All I know is that i is an integer")
	}
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b

}

func test4() {

	x := 3
	y := 4
	z := 5

	max_xy := max(x, y) //调用函数max(x, y)
	max_xz := max(x, z) //调用函数max(x, z)

	fmt.Printf("max(%d, %d) = %d\n", x, y, max_xy)
	fmt.Printf("max(%d, %d) = %d\n", x, z, max_xz)
	fmt.Printf("max(%d, %d) = %d\n", y, z, max(y, z)) // 也可在这直接调用它

}

//返回 A+B 和A*B
func SumAndProduct(a, b int) (int, int) {
	return a + b, a * b
}

func test5() {
	x := 3
	y := 4
	xPlus, xT := SumAndProduct(x, y)
	fmt.Printf("%d + %d = %d\n", x, y, xPlus)
	fmt.Printf("%d * %d = %d\n", x, y, xT)
	fmt.Println("type:", reflect.TypeOf(x)) //查看数据类型

}
