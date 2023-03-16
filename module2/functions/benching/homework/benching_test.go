package main

import "testing"

func BenchmarkMapUserProducts(b *testing.B) {

	users := genUsers()
	products := genProducts()

	for i := 0; i < b.N; i++ {
		MapUserProducts(users, products)
	}
}

func BenchmarkMapUserProducts2(b *testing.B) {

	users := genUsers()
	products := genProducts()

	for i := 0; i < b.N; i++ {
		MapUserProducts2(users, products)
	}
}
