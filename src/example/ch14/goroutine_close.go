/*
Project:ch14
Time:2016/9/7-16:02
Author:diogo
*/
// Q20_goroutine.go
package main

import (
	"fmt"
)

func tel(ch chan int) {
	for i:=0; i < 15; i++ {
		ch <- i
	}
	close(ch) // if this is ommitted: panic: all goroutines are asleep - deadlock!
}

func main() {
	var ok = true
	var i int
	ch := make(chan int,20)

	go tel(ch)
	for ok {
		if i, ok= <-ch; ok {
			fmt.Printf("ok is %t and the counter is at %d\n", ok, i)
		}
	}
}
