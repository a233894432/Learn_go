/*
Project:ch14
Time:2016/9/7-16:05
Author:diogo
*/
package main

import (
	"fmt"
	"os"
)

func tel(ch chan int, quit chan bool) {
	for i:=0; i < 15; i++ {
		ch <- i
	}
	quit <- true
}

func main() {
	var ok = true
	ch := make(chan int)
	quit := make(chan bool)

	go tel(ch, quit)
	for ok {
		select {
		case i:= <-ch:
			fmt.Printf("The counter is at %d\n", i)
		case <-quit:
			os.Exit(0)
		}
	}
}