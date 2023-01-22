package main

import "fmt"

func main() {
	var numberInt int
	numberInt = 3
	var numberFloat float32 = float32(numberInt)
	fmt.Printf("тип: %T, значение: %v", numberFloat, numberFloat)
}
