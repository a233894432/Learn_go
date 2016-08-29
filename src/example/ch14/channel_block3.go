/*
    author: diogoxiang
    time: 2016-08-16/8/30 - 01:48
    name: Learn_go
 */
package main

import "fmt"
import "time"

func main() {
	c := make(chan int)
	go func() {
		time.Sleep(15 * 1e9)
		x := <-c
		fmt.Println("received", x)
	}()
	fmt.Println("sending", 10)
	c <- 10
	fmt.Println("sent", 10)
}
/* Output:
sending 10
(15 s later):
received 10
sent 10
*/
