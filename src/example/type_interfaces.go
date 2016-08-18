/*
Project:Learn_go
Time:2016/8/18-17:46
author:Dio
*/
package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func main() {
	var areaIntf Shaper
	sq1 := new(Square)
	sq1.side = 5

	areaIntf = sq1
	// Is Square the type of areaIntf?
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

/*
输出：

The type of areaIntf is: *main.Square
areaIntf does not contain a variable of type Circle
程序行中定义了一个新类型 Circle，它也实现了 Shaper 接口。
t, ok := areaIntf.(*Square); ok 测试 areaIntf 里是否一个包含 'Square' 类型的变量，结果是确定的；
然后我们测试它是否包含一个 'Circle' 类型的变量，结果是否定的。

备注

如果忽略 areaIntf.(*Square) 中的 * 号，
会导致编译错误：impossible type assertion: Square does not implement Shaper (Area method has pointer receiver)。
 */