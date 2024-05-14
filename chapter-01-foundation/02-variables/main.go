package main

const text = "Hello world!"

var (
	flag        bool    = true
	number      int     = 10
	dummyString string  = "dummy string"
	decimal     float64 = 1.2
)

func main() {

	dynamicInference := "dynamic Inference variable type"
	println(text, flag, number, dummyString, decimal, dynamicInference)
}
