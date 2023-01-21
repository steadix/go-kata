package main

import (
	"fmt"
)

func main() {
	typeInt()
}

func typeInt() {
	fmt.Println("=== START type int ===")
	var uintNumber8 uint = 1 << 7
	var from8 = int8(uintNumber8)
	uintNumber8--
	var to8 = int8(uintNumber8)
	fmt.Println("min int8 number :", from8, "		       | max int8 number :", to8)
	var uintNumber16 uint = 1 << 15
	var from16 = int16(uintNumber16)
	uintNumber16--
	var to16 = int16(uintNumber16)
	fmt.Println("min int16 number:", from16, "	       | max int16 number:", to16)
	var uintNumber32 uint = 1 << 31
	var from32 = int32(uintNumber32)
	uintNumber32--
	var to32 = int32(uintNumber32)
	fmt.Println("min int32 number:", from32, "	       | max int32 number:", to32)
	var uintNumber64 uint = 1 << 63
	var from64 = int64(uintNumber64)
	uintNumber64--
	var to64 = int64(uintNumber64)
	fmt.Println("min int64 number:", from64, "| max int64 number:", to64)
	fmt.Println("=== END type int ===")

}
