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

	templates := []string{"header.html", "content.html", "footer.html"}

	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(os.Stdout, Courses{
		{Name: "Go Programming", Workload: 40},
		{Name: "Docker", Workload: 20},
		{Name: "Kubernetes", Workload: 30},
	})
	if err != nil {
		panic(err)
	}

}
