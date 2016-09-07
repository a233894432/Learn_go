/*
   author: diogoxiang
   time: 2016-09-16/9/7 - 00:00
   name: Learn_go
*/
package main

import "fmt"

func sum(x, y int, ch chan int) {

	ch <- x + y
}

func main() {
	c := make(chan int)
	go sum(12, 13, c)
	fmt.Println(<-c)
}
