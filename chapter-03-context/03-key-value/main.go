package main

import (
	"context"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	ctx := context.Background()
	ctx = context.WithValue(ctx, "token", "hash-token")
	ctx = context.WithValue(ctx, "person", Person{Name: "John Doe", Age: 30})
	BookHotel(ctx)
}

func BookHotel(ctx context.Context) {
	token := ctx.Value("token").(string)
	person := ctx.Value("person").(Person)
	fmt.Println("Token:", token)
	fmt.Println("Person:", person.Name, person.Age)

}
