/*
Project:Learn_go
Time:2016/8/19-15:36
author:Dio
*/
package main

import "fmt"

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 1; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {
	// ...
	Name string
}

func (b *Bird) Quack() {
	fmt.Println("I am quacking!",b.Name)
}

func (b *Bird) Walk()  {
	fmt.Println("I am walking!")
}

func main() {
	b := new(Bird)
	b.Name="diogo";
	DuckDance(b)
}