/*
Project:Learn_go
Time:2016/8/18-14:21
author:Dio
*/
/*
多重继承
*/
package main

import (
	"fmt"
)

type Base struct{}

func (Base) Magic() {
	fmt.Println("base magic")
}

func (self Base) MoreMagic() {
	self.Magic()
	self.Magic()
}

type Voodoo struct {
	Base
}

func (Voodoo) Magic() {
	fmt.Println("voodoo magic")
}

func main() {
	v := new(Voodoo)
	v.Magic()
	v.MoreMagic()
	/*
	voodoo magic
	base magic
	base magic
	*/
}
