package main

import (
	"fmt"
	"unsafe"
)

func main() {
	n := 112358132134
	fmt.Println("size is:", unsafe.Sizeof(n), "bytes")
	var numberUint8 uint8 = 1 >> 1
	fmt.Println("left shift uint8:", numberUint8)

}
