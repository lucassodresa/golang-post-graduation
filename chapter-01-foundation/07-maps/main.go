package main

import "fmt"

func main() {

	salaries := map[string]int{"Lucas": 1000, "John": 2000, "Doe": 3000}
	delete(salaries, "Lucas")
	salaries["Luc"] = 5000

	// salariesWithMake := make(map[string]int)
	// salariesWithoutMake := map[string]int{}

	for name, salary := range salaries {
		fmt.Printf("%s earns %d\n", name, salary)
	}

	for _, salary := range salaries {
		fmt.Printf("Salary is %d\n", salary)
	}

}
