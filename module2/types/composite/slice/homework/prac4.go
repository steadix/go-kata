package main

import "fmt"

type User3 struct {
	Name string
	Age  int
}

func main() {
	users := []User3{
		{
			Name: "Eva",
			Age:  13,
		},
		{
			Name: "Victor",
			Age:  28,
		},
		{
			Name: "Dex",
			Age:  34,
		},
		{
			Name: "Billy",
			Age:  21,
		},
		{
			Name: "Foster",
			Age:  29,
		},
	}

	subUsers := make([]User3, len(users))
	copy(subUsers, users)
	subUsers = subUsers[2:]
	editSecondSlice(subUsers)
	fmt.Println(users)
}

func editSecondSlice(users []User3) {
	for i := range users {
		users[i].Name = "unknown"
	}
}
