package main

import (
	"fmt"
	"unsafe"
)

func main() {
	n := 112358132134
	fmt.Println(unsafe.Sizeof(n))
}
