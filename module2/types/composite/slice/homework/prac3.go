package main

import "fmt"

type User2 struct {
	Name string
	Age  int
}

func main() {
	users := []User2{
		{
			Name: "Alice",
			Age:  21,
		},
		{
			Name: "John",
			Age:  34,
		},
		{
			Name: "Alexander",
			Age:  45,
		},
		{
			Name: "Ivan",
			Age:  13,
		},
		{
			Name: "Denis",
			Age:  44,
		},
		{
			Name: "Mary",
			Age:  26,
		},
		{
			Name: "Rose",
			Age:  41,
		},
	}
	_, users = users[len(users)-1], users[:len(users)-1]
	_, users = users[0], users[1:]
	fmt.Println(users)
}
