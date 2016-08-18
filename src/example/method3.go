/*
Project:Learn_go
Time:2016/8/18-13:52
author:Dio
*/
/*
当一个匿名类型被内嵌在结构体中时，匿名类型的可见方法也同样被内嵌，
这在效果上等同于外层类型 继承 了这些方法：将父类型放在子类型中来实现亚型。
这个机制提供了一种简单的方式来模拟经典面向对象语言中的子类和继承相关的效果，
也类似 Ruby 中的混入（mixin）。
 */
package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

type NamedPoint struct {
	Point
	name string
}
func (n *NamedPoint) Abs() float64 {
	return n.Point.Abs() * 100.
}

func main() {
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}
	fmt.Println(n.Abs()) // 打印500
}
