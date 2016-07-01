package main

import (
	"fmt"
	"time"
)

func main() {

	//最基本的类型，一个单一的循环
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	fmt.Println(time.Now())
	//一个精典的类型：带初始化/条件的循环
	for j := 7; j <= 88; j++ {
		fmt.Println(j)
	}

	//通过条件的for循环会一直执行，直到你中断或在闭包中返回。
	for {
		fmt.Println("loop")
		break
	}

	//你可以使用多种条件匹配，同样你可以使用默认匹配
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("it's the weekend")
	default:
		fmt.Println("it's a weekday")
	}

	fmt.Println(time.Now())

}
