/*
    author: diogoxiang
    time: 2016-08-16/8/22 - 22:19
    name: Learn_go
 */
package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("src/example/products.txt")
	if err != nil && os.IsNotExist(err) {
		fmt.Printf("file not exist!\n")
		return
	}
	fmt.Printf("file exist!\n")
	defer f.Close()
}
