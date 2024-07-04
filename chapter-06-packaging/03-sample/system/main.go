package main

import (
	"github.com/google/uuid"
	"github.com/lucassodresa/golang-post-graduation/chapter-06-packaging/03-sample/math"
)

func main() {
	m := math.NewMath(1, 2)
	println(m.Add())

	println(uuid.New().String())
}
