package main

import (
	"fmt"
	"github.com/brianvoe/gofakeit/v6"
	"math/rand"
)

type User struct {
	ID       int64
	Name     string `fake:"{firstname}"`
	Products []Product
}

type Product struct {
	UserID int64
	Name   string `fake:"{sentence:3}"`
}

func main() {
	users := genUsers()
	fmt.Println(users)
	products := genProducts()
	fmt.Println(products)
	/* users = MapUserProducts(users, products)
	fmt.Println(users) */
	users = MapUserProducts2(users, products)
	fmt.Println(users)
}

func MapUserProducts(users []User, products []Product) []User {

	for i, user := range users {
		for _, product := range products {
			if product.UserID == user.ID {
				users[i].Products = append(users[i].Products, product)
			}
		}
	}

	return users
}

func MapUserProducts2(users []User, products []Product) []User {
	productsMap := make(map[int64]*User)

	for i := range users {
		productsMap[users[i].ID] = &users[i]
	}

	for i := range products {
		userID := products[i].UserID
		if user, ok := productsMap[userID]; ok {
			user.Products = append(user.Products, products[i])
		}
	}

	return users
}

func genProducts() []Product {
	products := make([]Product, 1000)
	for i, product := range products {
		_ = gofakeit.Struct(&product)
		product.UserID = int64(rand.Intn(100) + 1)
		products[i] = product
	}

	return products
}

func genUsers() []User {
	users := make([]User, 100)
	for i, user := range users {
		_ = gofakeit.Struct(&user)
		user.ID = int64(i)
		users[i] = user
	}

	return users
}
