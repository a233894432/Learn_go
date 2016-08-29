/*
   author: diogoxiang
   time: 2016-08-16/8/30 - 01:40
   name: Learn_go
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	time.Sleep(10 * 1e9)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(1e9)
	for {
		input = <-ch
		fmt.Printf("%s ", input)
	}
}
