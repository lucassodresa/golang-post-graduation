package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

func main() {
	course := Course{Name: "Go Programming", Workload: 40}
	tmp := template.New("CourseTemplate")
	tmp, _ = tmp.Parse("Course Name: {{.Name}}\nWorkload: {{.Workload}} hours\n")
	err := tmp.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
