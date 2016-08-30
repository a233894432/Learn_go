/*
Project:Learn_go
Time:2016/8/17-15:04
author:Dio
*/
/**
	new
	make
	区别..
	不能编译
	试图 make() 一个结构体变量，会引发一个编译错误，这还不是太糟糕，
	但是 new() 一个映射并试图使用数据填充它，将会引发运行时错误！
	因为 new(Foo) 返回的是一个指向 nil 的指针，它尚未被分配内存。所以在使用 map 时要特别谨慎。
 */
package main

type Foo map[string]string
type Bar struct {
	thingOne string
	thingTwo int
}

func main() {
	// OK
	y := new(Bar)
	(*y).thingOne = "hello"
	(*y).thingTwo = 1

	// NOT OK
	//z := make(Bar) // 编译错误：cannot make type Bar
	//(*z).thingOne = "hello"
	//(*z).thingTwo = 1

	// OK
	x := make(Foo)
	x["x"] = "goodbye"
	x["y"] = "world"

	// NOT OK
	u := new(Foo)
	(*u)["x"] = "goodbye" // 运行时错误!! panic: assignment to entry in nil map
	(*u)["y"] = "world"
}
