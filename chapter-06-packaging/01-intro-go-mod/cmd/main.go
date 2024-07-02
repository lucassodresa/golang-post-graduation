package main

import (
	"fmt"

	"github.com/lucassodresa/golang-post-graduation/chapter-06-packaging/01-intro-go-mod/math"
)

func main() {
	m := math.Math{A: 3, B: 4}
	fmt.Println(m.Add())
	// fmt.Println("Hello, World!")
}
