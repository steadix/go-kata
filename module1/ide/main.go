package main

import (
	"fmt"
	"math/rand"
	"os"
)

const (
	noError = iota
	generalError
)

func printme(text string) {
	fmt.Println(text)
}

func main() {
	printme("Hello linter")
	roll := rand.Intn(100)
	if roll > 50 {
		os.Exit(noError)
	} else {
		os.Exit(generalError)
	}
}
