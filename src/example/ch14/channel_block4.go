/*
    author: diogoxiang
    time: 2016-09-16/9/6 - 23:44
    name: Learn_go
 */
package main

import (
	"time"
	"fmt"
)

func printer(c chan int) {
	for i := 0; i < 10; i++ {
		c <- i
		time.Sleep(time.Second)
	}
	close(c)
}

func reciever(c chan int) {
	for recievedMsg := range c {
		fmt.Println(recievedMsg)
	}
}

func main() {
	newChanel := make(chan int)
	go printer(newChanel)
	reciever(newChanel)
}
