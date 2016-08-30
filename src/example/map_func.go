/*
Project:Learn_go
Time:2016/8/16-17:45
author:Dio
*/
package main
import "fmt"

func main() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf[1])
}