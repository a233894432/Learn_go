package main

import (
	"github.com/kataras/iris"
)

type Person struct {
	Name   string
	Age    int
	Emails []string
	Jobs   []*Job
	
}

type Job struct {
	Employer string
	Role     string
}

func main() {
	iris.Get("/hi", func(ctx *iris.Context) {
		ctx.Write("Hi %s", "iris")
	})
	iris.Listen(":8080")
	//	err := iris.ListenWithErr(":8080")
}
