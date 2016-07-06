package main

import (
	"fmt"
)

func main() {
LABEL1:
	for i := 1; i < 10; i++ {
		for {
			fmt.Println(i)
			continue LABEL1
		}
	}

	fmt.Println("ok")

	fmt.Println("okssss")

}
