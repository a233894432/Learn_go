/*
Project:myexample
Time:2016年10月12日10:38:54
Author:diogo
*/
package main

func main() {
	println("----")
	utest(21, 3)

}

func test() {
	x, y := 10, 20

	defer func(i int) {
		println("defer:", i, y) // y 闭包引用
	}(x) // x 被复制

	x += 10
	y += 100
	println("x =", x, "y =", y)

	/*
		x = 20 y = 120
		defer: 10 120
	*/
}

func utest(x, y int) {
	var z int

	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()

		z = x / y
		return
	}()

	println("x / y =", z)
}
