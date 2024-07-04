package main

import (
	"fmt"

	"github.com/lucassodresa/golang-post-graduation/chapter-06-packaging/01-intro-go-mod/math"
)

func main() {
	m := math.NewMath(1, 2)
	m.C = 3
	fmt.Println(m.Add())
	fmt.Println(m.C)
}
