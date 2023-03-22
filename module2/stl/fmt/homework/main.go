package main

import (
	"fmt"
)

type Person struct {
	Name  string
	Age   int
	Money float64
}

func generateSelfStory(p Person) string {
	return fmt.Sprintf("Hello! My name is %s. I'm %d y.o. And I also have $%.2f in my pocket right now.", p.Name, p.Age, p.Money)
}

func main() {
	p := Person{Name: "Andy", Age: 18, Money: 10.000025}

	story := generateSelfStory(p)
	fmt.Println(story)
}
