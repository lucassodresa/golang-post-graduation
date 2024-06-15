package main

import (
	"os"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Courses{
		{Name: "Go Programming", Workload: 40},
		{Name: "Docker", Workload: 20},
		{Name: "Kubernetes", Workload: 30},
	})
	if err != nil {
		panic(err)
	}
}
