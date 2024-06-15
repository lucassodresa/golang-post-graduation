package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	Name     string
	Workload int
}

type Courses []Course

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.New("template.html").ParseFiles("template.html"))
		err := t.Execute(w, Courses{
			{Name: "Go Programming", Workload: 40},
			{Name: "Docker", Workload: 20},
			{Name: "Kubernetes", Workload: 30},
		})
		if err != nil {
			panic(err)
		}
	})

	http.ListenAndServe(":8282", nil)

}
