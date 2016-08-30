//2016/8/12 15:40
/**
使用 defer 语句来记录函数的参数与返回值

下面的代码展示了另一种在调试时使用 defer 语句的手法（示例 6.12 defer_logvalues.go）：
 */
package main

import (
	"io"
	"log"
)

func func1(s string) (n int, err error) {
	defer func() {
		log.Printf("func1(%q) = %d, %v", s, n, err)
	}()
	return 7, io.EOF
}

func main() {
	func1("Go")
}

/* out:
	2016/08/12 15:41:28 func1("Go") = 7, EOF
 */