package main

import (
	"fmt"
	"unsafe"
)

func main() {
	typeFloat()
}

func typeFloat() {
	fmt.Println("=== START type float ===")
	var uintNumber uint32 = 1 << 29
	uintNumber += 1 << 28
	uintNumber += 1 << 27
	uintNumber += 1 << 26
	uintNumber += 1 << 25
	uintNumber += 1 << 21
	uintNumber += 1 << 31
	var floatNumber float32
	floatNumber = *(*float32)(unsafe.Pointer(&uintNumber))
	fmt.Println(floatNumber)
	fmt.Println("=== END type float ===")
}
