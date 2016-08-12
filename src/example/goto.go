//2016/8/12 9:41
package main

import "fmt"

func main() {
	i := 0
HERE:
	fmt.Println(i)
	i++
	if i == 5 {
		return
	}
	goto HERE
}
