//2016/8/11 14:23
package main

func main() {
	/**
		创建一个程序，要求能够打印类似下面的结果（直到每行 25 个字符时为止）：

	G
	GG
	GGG
	GGGG
	GGGGG
	GGGGGG
	使用 2 层嵌套 for 循环。
	使用一层 for 循环以及字符串截断。

	*/
	// 1 - use 2 nested for loops
	for i := 1; i <= 25; i++ {
		for j := 1; j <= i; j++ {
			print("G")
		}
		println()
	}
	// 2 -  use only one for loop and string concatenation
	str := "G"
	for i := 1; i <= 25; i++ {
		println(str)
		str += "G"
	}
}
