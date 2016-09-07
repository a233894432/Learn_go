/*
Project:ch14
Time:2016/9/7-15:57
Author:diogo
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)

	go pump1(ch1)
	go pump2(ch2)
	go suck(ch1, ch2)

	time.Sleep(1e9)

	/** out
		输出：

	Received on channel 2: 5
	Received on channel 2: 6
	Received on channel 1: 0
	Received on channel 2: 7
	Received on channel 2: 8
	Received on channel 2: 9
	Received on channel 2: 10
	Received on channel 1: 2
	Received on channel 2: 11
	...
	Received on channel 2: 47404
	Received on channel 1: 94346
	Received on channel 1: 94348

	*/
}

func pump1(ch chan int) {
	for i := 0; ; i++ {
		ch <- i * 2
	}
}

func pump2(ch chan int) {
	for i := 0; ; i++ {
		ch <- i + 5
	}
}

func suck(ch1, ch2 chan int) {
	for {
		select {
		case v := <-ch1:
			fmt.Printf("Received on channel 1: %d\n", v)
		case v := <-ch2:
			fmt.Printf("Received on channel 2: %d\n", v)
		}
	}
}
