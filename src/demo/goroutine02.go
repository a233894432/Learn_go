package main

import "fmt"

func sum(a []int, c chan int) {
	total := 0
	for _, v := range a {
		total += v
	}
	c <- total // send total to c
}

func main() {
	a := []int{7, 2, 8, -9, 4, 0, 12, 13}
	d := len(a)
	c := make(chan int)
	fmt.Println(a[:d/2])
	fmt.Println(a[d/2:])
	go sum(a[:d/2], c)
	x := <-c
	go sum(a[d/2:], c)
	y := <-c
	//	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}
