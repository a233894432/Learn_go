/*
Project:myexample
Time:2016/9/30-11:33
Author:diogo
*/
package main

import "fmt"

func main() {

	s := []int{1, 2, 3, 4, 5}

	for i, v := range s { // 复制 struct slice { pointer, len, cap }。

		if i == 0 {
			s = s[:3]  // 对 slice 的修改，不会影响 range。
			s[2] = 100 // 对底层数据的修改。
		}

		fmt.Println(i, v)
	}

	/*out
	0 1
	1 2
	2 100
	3 4
	4 5
	*/

	x := []int{1, 2, 3}
	i := 2

	switch i {
	case x[1]:
		println("a")

	case 1, 3:
		println("b")

	default:
		println("c")
	}

	switch {
	case x[1] > 0:
		println("a")
	case x[1] < 0:
		println("b")
	default:
		println("c")
	}

	switch i := x[2]; { // 带初始化语句
	case i > 0:
		println("a")
	case i < 0:
		println("b")
	default:
		println("c")
	}
}
