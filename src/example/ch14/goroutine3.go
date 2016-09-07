/*
<<<<<<< HEAD
   author: diogoxiang
   time: 2016-09-16/9/7 - 00:03
   name: Learn_go
=======
Project:ch14
Time:2016/9/7-15:24
Author:diogo
>>>>>>> origin/learn_web_application
*/
package main

import "fmt"

<<<<<<< HEAD
/**
有两个协程，第一个提供数字 0，10，20，...90 并将他们放入通道，第二个协程从通道中读取并打印。
main() 等待两个协程完成后再结束。
 */

// integer producer:
func numGen(start, count int, out chan<- int) {
	for i := 0; i < count; i++ {
		out <- start
		start = start + count
	}
	close(out)
}

// integer consumer:
func numEchoRange(in <-chan int, done chan<- bool) {
	for num := range in {
		fmt.Printf("%d\n", num)
	}
	done <- true
}

func main() {
	numChan := make(chan int)
	done := make(chan bool)
	go numGen(10, 10, numChan)
	go numEchoRange(numChan, done)

	<-done
}

/* Output:
0
10
20
30
40
50
60
70
80
90
*/
=======
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
>>>>>>> origin/learn_web_application
