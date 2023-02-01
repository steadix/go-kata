package main

import (
	"fmt"
)

type MyInterface interface{}

func main() {
	var n *int
	fmt.Println(n == nil)
	test(n)
}

func test(r interface{}) {
	switch v := r.(type) {
	case *int:
		if v == nil {
			fmt.Println("Success!")
		}
	}
}
