/*
Project:example
Time:2016/8/29-11:11
Author:diogo
*/
package main

import (
	"fmt"
	"io"
	"os"
)
func main() {
	CopyFile("./example/products.txt", "./example/source.txt")
	fmt.Println("Copy done!")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
