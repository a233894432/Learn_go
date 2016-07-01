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
	iris.Config.Render.Template.Extensions = []string{".jade"}
	// this is optionally, you can keep .html extension
	iris.Config.Render.Template.Engine = iris.JadeEngine

	iris.Get("/", func(ctx *iris.Context) {

		job1 := Job{Employer: "Super Employer", Role: "Team leader"}
		job2 := Job{Employer: "Fast Employer", Role: "Project managment"}

		person := Person{
			Name:   "name1",
			Age:    50,
			Emails: []string{"email1@something.gr", "email2.anything@gmail.com"},
			Jobs:   []*Job{&job1, &job2},
		}

		ctx.MustRender("page.jade", person)
	})

	iris.Listen(":8080")
}
