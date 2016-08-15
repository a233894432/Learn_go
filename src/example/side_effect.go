//2016/8/12 14:21
package main

import (
	"fmt"
)

// this function changes reply:
func Multiply(a, b int, reply *int) {
	*reply = a * b
}

func main() {
	n := 0
	reply := &n
	fmt.Println("初始值: n===",n) //初始值  n  === 0
	fmt.Printf("指针:reply= &n== %p",reply)  //指针: reply= &n== 0xc082002278

	fmt.Println()

	Multiply(10, 5, reply)

	fmt.Println("Multiply:", *reply) // Multiply: 50
	fmt.Printf("改就后指针:reply= &n== %p",reply)
}