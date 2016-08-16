/*
Project:Learn_go
Time:2016/8/16-15:20
author:Dio
*/
package main

import "fmt"

func main() {

	slice1 := make([]int, 0, 10)
	// load the slice, cap(slice1) is 10:
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	// print the slice:
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	var ar = [10]int{0,1,2,3,4,5,6,7,8,9}
	fmt.Println(ar[5:7])
}
