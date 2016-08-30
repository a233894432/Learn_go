/*
   author: diogoxiang
   time: 2016-08-16/8/26 - 23:50
   name: Learn_go
*/
package main

import (
	"fmt"
	"os"
)

var user = os.Getenv("USER")

func check() {
	if user == "" {
		panic("Unknown user: no value for $USER")
	}
}
func init() {
	check()
}

func main() {
	fmt.Println("Starting the program")
	panic("A severe error occurred: stopping the program!")
	fmt.Println("Ending the program")
}
