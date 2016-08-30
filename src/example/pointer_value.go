/*
Project:Learn_go
Time:2016/8/17-18:21
author:Dio
*/
package main

import (
	"fmt"
	"math"
	"time"
)

type B struct {
	thing int
}

func (b *B) change() { b.thing = 1 }

func (b B) write() string { return fmt.Sprint(b) }

type Point3 struct{ x, y, z float64 }

// A method on Point3
func (p Point3) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y + p.z*p.z)
}

func main() {
	var b1 B // b1是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(B) // b2是指针
	b2.change()
	fmt.Println(b2.write())



	start := time.Now()
	//值拷贝
	p3:=Point3{3,4,5}
	p3.Abs();
	end := time.Now()
	delta := end.Sub(start)

	fmt.Println("addTime====", delta)



}
