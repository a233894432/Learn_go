/*
Project:Learn_go
Time:2016/8/18-16:15
author:Dio
*/
package main

import "fmt"

type Shaper interface {
	Area() float32
}

type Square struct {
	side float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {

	r := &Rectangle{5, 3} // Area() of Rectangle needs a value
	q := &Square{5}      // Area() of Square needs a pointer
	// shapes := []Shaper{Shaper(r), Shaper(q)}
	// or shorter
	shapes := []Shaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	for n, _ := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}
}

/*
输出：

Looping through shapes for area ...
Shape details:  {5 3}
Area of this shape is:  15
Shape details:  &{5}
Area of this shape is:  25


在调用 shapes[n].Area() 这个时，只知道 shapes[n] 是一个 Shaper 对象，
最后它摇身一变成为了一个 Square 或 Rectangle 对象，并且表现出了相对应的行为。

也许从现在开始你将看到通过接口如何产生 更干净、更简单 及 更具有扩展性 的代码。
 */