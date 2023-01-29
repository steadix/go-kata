package main

import (
	"fmt"
	"unsafe"
)

type User struct {
	Age      int
	Name     string
	Wallet   Wallet
	Location Location
}
type Location struct {
	Address string
	City    string
	index   string
}
type Wallet struct {
	USD uint64
	RUR uint64
	BTC uint64
	ETH uint64
}

func main() {
	user := User{
		Age:  13,
		Name: "Alexander",
	}
	fmt.Println(user)
	wallet := Wallet{
		USD: 250000,
		RUR: 3500,
		BTC: 1,
		ETH: 4,
	}
	user2 := User{
		Age:  34,
		Name: "Anton",
		Wallet: Wallet{
			USD: 144000,
			RUR: 8900,
			BTC: 55,
			ETH: 34,
		},
		Location: Location{
			Address: "Нововатутинская 3-Я ул, 13, к.2",
			City:    "Москва",
			index:   "108836",
		},
	}
	fmt.Println(wallet)
	fmt.Println("wallet allocates: ", unsafe.Sizeof(wallet), "bytes")
	user.Wallet = wallet
	fmt.Println(user)
	fmt.Println(user2)
}
