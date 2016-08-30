//2016/8/12 17:18
package main

import "fmt"

func main() {
	f()
}
func f() {
	for i := 0; i < 4; i++ {
		g := func(i int) { fmt.Printf("%d ", i) }
		//此例子中只是为了演示匿名函数可分配不同的内存地址，在现实开发中，不应该把该部分信息放置到循环中。
		g(i)
		fmt.Printf(" - g is of type %T and has value %v\n", g, g)
	}
}

/* out:
0  - g is of type func(int) and has value 0x401210
1  - g is of type func(int) and has value 0x401210
2  - g is of type func(int) and has value 0x401210
3  - g is of type func(int) and has value 0x401210
*/
