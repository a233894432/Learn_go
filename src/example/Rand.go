/*
   author: diogoxiang
   time: 2016-08-16/8/16 - 23:40
   name: Learn_go
*/
/*	代码片段:
	生成随机数
*/
package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 10; i++ {
		fmt.Println(r.Intn(100))
	}
}
