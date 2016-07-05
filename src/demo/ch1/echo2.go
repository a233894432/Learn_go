package main

import (
	"fmt"
)

func main() {

	for i := 1; i < 10; i++ {
		for {
			fmt.Println(i)
			goto LABEL1
		}
	}
LABEL1:
	fmt.Println("ok")

}
