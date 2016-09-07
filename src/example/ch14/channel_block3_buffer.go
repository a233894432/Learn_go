/*
    author: diogoxiang
    time: 2016-09-16/9/6 - 23:58
    name: Learn_go
 */
package main

import "fmt"
import "time"

func main() {
	c := make(chan int, 50)
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
sent 10   // prints immediately
no further output, because main() then stops
*/