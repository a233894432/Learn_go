/*
Project:Learn_go
Time:2016/8/18-14:07
author:Dio
*/
/*
A：聚合（或组合）：包含一个所需功能类型的具名字段。

B：内嵌：内嵌（匿名地）所需功能类型，像前一节 10.6.5 所演示的那样。

为了使这些概念具体化，假设有一个 Customer 类型，我们想让它通过 Log 类型来包含日志功能，
Log 类型只是简单地包含一个累积的消息（当然它可以是复杂的）。
如果想让特定类型都具备日志功能，你可以实现一个这样的 Log 类型，然后将它作为特定类型的一个字段，并提供 Log()，它返回这个日志的引用。

方式 A 可以通过如下方法实现（使用了第 10.7 节中的 String() 功能）：
 */
package main

import (
	"fmt"
)

type Log struct {
	msg string
}

type Customer struct {
	Name string
	log  *Log
}

func main() {
	//long A
	c := new(Customer)
	c.Name = "Barak Obama"
	c.log = new(Log)
	c.log.msg = "1 - Yes we can!"
	// shorter B
	// c = &Customer{"Barak Obama", &Log{"2 - Yes we can!"}}

	// fmt.Println(c) &{Barak Obama 1 - Yes we can!}
	//c.Log().Add("2 - After me the world will be a better place!")
	//c.Log().Add("3 - 333place!")
	//c.Log().Add("4 - 4444 place!")
	//fmt.Println(c.log)
	fmt.Println(c.Log())
	fmt.Println(c.log.String())

}

func (l *Log) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log) String() string {
	return l.msg
}

func (c *Customer) Log() *Log {
	return c.log
}
