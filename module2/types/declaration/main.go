package main

import (
	"fmt"
	"os"
)

const (
	NoError = iota
	GeneralError
	InternalError
)

const (
	NoErrorMsg       = "no error"
	GeneralErrorMsg  = "general error"
	InternalErrorMsg = "internal error"
)

var ErrorDic = map[int]string{
	NoError:       NoErrorMsg,
	GeneralError:  GeneralErrorMsg,
	InternalError: InternalErrorMsg,
}

var (
	GlobalVar1 float32 = 55.89
	GlobalVar2         = "some message"
)

func main() {
	n := 13
	text := "some text from Kata"
	fmt.Println(n, text)
	n2, text2, fl := 21, "one line declaration", 3.14
	fmt.Println(n2, text2, fl)
	var n3 int
	var n4 int64
	var fl2 float32
	n3, n4, fl2 = 3, 5, 8.13
	fmt.Println(n3, n4, fl2)
	var fl3 float32 = 6.63
	fmt.Println(fl3)
	fl4 := float32(1.05)
	_ = fl4

	var (
		id, age        int
		name, nickname string
		balance        float64
	)
	_, _, _, _, _ = id, age, name, nickname, balance
	_, _ = GlobalVar1, GlobalVar2
	fmt.Println(ErrorDic[NoError])
	os.Exit(NoError)

}
