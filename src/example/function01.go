/*
   author: diogoxiang
   time: 2016-08-16/8/15 - 23:21
   name: Learn_go
*/

/*
练习 6.10

学习并理解以下程序的工作原理：

一个返回值为另一个函数的函数可以被称之为工厂函数，这在您需要创建一系列相似的函数的时候非常有用：
书写一个工厂函数而不是针对每种情况都书写一个函数。下面的函数演示了如何动态返回追加后缀的函数：

func MakeAddSuffix(suffix string) func(string) string {
    return func(name string) string {
        if !strings.HasSuffix(name, suffix) {
            return name + suffix
        }
        return name
    }
}
现在，我们可以生成如下函数：

addBmp := MakeAddSuffix(“.bmp”)
addJpeg := MakeAddSuffix(“.jpeg”)
然后调用它们：

addBmp("file") // returns: file.bmp
addJpeg("file") // returns: file.jpeg
可以返回其它函数的函数和接受其它函数作为参数的函数均被称之为高阶函数，是函数式语言的特点。
我们已经在第 6.7 中得知函数也是一种值，因此很显然 Go 语言具有一些函数式语言的特性。
闭包在 Go 语言中非常常见，常用于 goroutine 和管道操作（详见第 14.8-14.9 节）。
在第 11.14 节的程序中，我们将会看到 Go 语言中的函数在处理混合对象时的强大能力。
*/
package main

import (
	"fmt"
	"strings"
)

func main() {
	addBmp := MakeAddSuffix(".bmp")
	addJpeg := MakeAddSuffix(".jpeg")
	fmt.Println(addBmp("file"))
	fmt.Println(addJpeg("diogoxiang"))
}

func MakeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}
