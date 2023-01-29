package main

import "fmt"

type User struct {
	Age    int
	Name   string
	Wallet Wallet
}

type Wallet struct {
	USD uint64
	RUR uint64
	BTC uint64
	ETH uint64
}

func main() {
	testArray()
	rangeArray()
}
func testArray() {
	a := [...]int{34, 55, 89, 144}
	fmt.Println("original value", a)
	a[0] = 21
	fmt.Println("changed first element", a)

	b := a

	a[0] = 233
	fmt.Println("original array", a)
	fmt.Println("modified array", b)

}

func rangeArray() {
	users := [4]User{
		{
			Age:  8,
			Name: "John",
			Wallet: Wallet{
				USD: 1,
				RUR: 13,
				BTC: 0,
				ETH: 0,
			},
		},
		{
			Age:  13,
			Name: "Katie",
			Wallet: Wallet{
				USD: 3,
				RUR: 500,
				BTC: 0,
				ETH: 0,
			},
		},
		{
			Age:  21,
			Name: "Doe",
			Wallet: Wallet{
				USD: 300,
				RUR: 0,
				BTC: 1,
				ETH: 3,
			},
		},
		{
			Age:  34,
			Name: "Arnie",
			Wallet: Wallet{
				USD: 34,
				RUR: 987342,
				BTC: 1,
				ETH: 3,
			},
		},
	}
	fmt.Println("Пользователи старше 18 лет:")
	for i := range users {
		if users[i].Age >= 18 {
			fmt.Println(users[i])
		}
	}
	fmt.Println("Пользователи, у которых есть криптовалюта на балансе:")
	for i, user := range users {
		if users[i].Wallet.BTC > 0 || users[i].Wallet.ETH > 0 {
			fmt.Println("user:", user, "index:", i)
		}
	}
	fmt.Println("========== end users non zero crypto balance ==========")

}
