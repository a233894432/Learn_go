/*
Project:myexample
Time:2016年10月11日17:05:00
Author:diogo
*/

/*

3.1 函数定义
不支持 嵌套 (nested)、重载 (overload) 和 默认参数 (default parameter)。

无需声明原型。
支持不定长变参。
支持多返回值。
支持命名返回参数。
支持匿名函数和闭包。

*/

package main

import "fmt"

func add(x, y int) (z int) {
	defer func() {
		println(z + 2) // 输出: 203
	}()

	z = x + y
	return z + 200 // 执行顺序: (z = z + 200) -> (call defer) -> (ret)
}

func main() {
	println(add(1, 2)) // 输出: 203
	f := test()        //x (0x2101ef018) = 100
	f()                //x (0x2101ef018) = 100

}

/*
3.4 匿名函数
匿名函数可赋值给变量，做为结构字段，或者在 channel 里传送。
// --- function variable ---

fn := func() { println("Hello, World!") }
fn()

// --- function collection ---
fns := [](func(x int) int){

    func(x int) int { return x + 1 },
    func(x int) int { return x + 2 },
}

println(fns[0](100))

// --- function as field ---

d := struct {
    fn func() string
}{
    fn: func() string { return "Hello, World!" },
}

println(d.fn())

// --- channel of function ---

fc := make(chan func() string, 2)
fc <- func() string { return "Hello, World!" }
println((<-fc)())
// ------function-----

   闭包:
   在汇编层面，test 实际返回的是 FuncVal 对象，其中包含了匿名函数地址、闭包对象指针。当调用匿名函数时，只需以某个寄存器传递该对象即可。

    FuncVal { func_address, closure_var_pointer ... }

*/
func test() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}
