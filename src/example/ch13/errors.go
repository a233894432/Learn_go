/*
   author: diogoxiang
   time: 2016-08-16/8/26 - 23:43
   name: Learn_go
*/
package main

import (
	"errors"
	"fmt"
)

var errNotFound error = errors.New("Not found error")

func main() {
	fmt.Printf("error: %v", errNotFound)
}
