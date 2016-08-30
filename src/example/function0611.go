/*
   author: diogoxiang
   time: 2016-08-16/8/15 - 23:37
   name: Learn_go
*/
/*
6.11 计算函数执行时间
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	start := time.Now()
	addTime()
	end := time.Now()
	delta := end.Sub(start)

	fmt.Println("addTime====", delta)

}

func addTime() {
	for i := 0; i < 20000000; i++ {

		v := i + 1
		fmt.Printf("%d  \n", v)
	}

}
