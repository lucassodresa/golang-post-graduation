package main

import (
	"html/template"
	"os"
	"strings"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {

	templates := []string{"header.html", "content.html", "footer.html"}

	t := template.New("content.html")
	t.Funcs(template.FuncMap{"ToUpper": strings.ToUpper})
	t = template.Must(t.ParseFiles(templates...))

	err := t.Execute(os.Stdout, Courses{
		{Name: "Go Programming", Workload: 40},
		{Name: "Docker", Workload: 20},
		{Name: "Kubernetes", Workload: 30},
	})
	if err != nil {
		panic(err)
	}

}
