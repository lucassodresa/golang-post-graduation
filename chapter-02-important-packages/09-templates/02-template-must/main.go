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
	t := template.Must(template.New("CourseTemplate").Parse("Course Name: {{.Name}}\nWorkload: {{.Workload}} hours\n"))
	err := t.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
