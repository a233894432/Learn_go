/*
Project:ch14
Time:2016/9/7-15:24
Author:diogo
*/
package main

import "fmt"

func main() {
	ch := make(chan string)
	go sendData(ch)
	getData(ch)
}

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokio"
	ch <- "Shenzhen"
	close(ch)
}

func getData(ch chan string) {
	for {
		input, open := <-ch
		if !open {
			fmt.Printf(" --- %v ",open)
			break
		}
		fmt.Printf("%s ", input)
	}
}
