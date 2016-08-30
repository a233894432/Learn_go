/*
    author: diogoxiang
    time: 2016-08-16/8/30 - 01:44
    name: Learn_go
 */
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	go pump(ch1)       // pump hangs
	go suck(ch1)
	fmt.Println(<-ch1) // prints only 0

	time.Sleep(1e9)
}

func pump(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}


func suck(ch chan int) {
	for {
		fmt.Println(<-ch)
	}
}