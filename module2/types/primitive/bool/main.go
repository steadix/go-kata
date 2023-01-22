package main

import (
	"fmt"
	"unsafe"
)

func main() {
	typeBool()
}

func typeBool() {
	var b bool
	fmt.Println("Размер в байтах:", unsafe.Sizeof(b))
	var u uint8 = 1
	fmt.Println(b)
	b = *(*bool)(unsafe.Pointer(&u))
	fmt.Println(b)
}
