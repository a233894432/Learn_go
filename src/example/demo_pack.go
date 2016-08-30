/*
Project:Learn_go
Time:2016/8/17-16:18
author:Dio
*/
package main

import (
	"fmt"
	"example/struct_pack"
)
func main() {
	struct1 := new(struct_pack.ExpStruct)
	struct1.Mi1 = 10
	struct1.Mf1 = 16.

	fmt.Printf("Mi1 = %d\n", struct1.Mi1)
	fmt.Printf("Mf1 = %f\n", struct1.Mf1)

	struct2:=new(struct_pack.ExpStruct2)

	struct2.Mf1=12
	struct2.Mi1=123
	fmt.Printf("Mi1 = %d\n", struct2.Mi1)
	fmt.Printf("Mf1 = %f\n", struct2.Mf1)



}