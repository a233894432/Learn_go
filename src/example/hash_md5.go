/*
Project:Learn_go
Time:2016/8/20-16:35
author:Dio
*/
package main

import (
	"fmt";
	"crypto/md5"
	"io"
)

func main() {
	hasher := md5.New()
	b := []byte{}
	io.WriteString(hasher, "diogo")
	fmt.Printf("Result: %x\n", hasher.Sum(b))
	fmt.Printf("Result: %d\n", hasher.Sum(b))

	date:=[]byte("diogo")
	fmt.Printf("%x ",md5.Sum(date))
}
/* Output:
Result: 098f6bcd4621d373cade4e832627b4f6
Result: [9 143 107 205 70 33 211 115 202 222 78 131 38 39 180 246]
*/