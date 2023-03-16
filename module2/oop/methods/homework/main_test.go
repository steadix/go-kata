package main

import (
	"testing"
)

func TestSum(t *testing.T) {

	type test struct {
		A    float64
		B    float64
		want float64
	}

	tests := []test{
		{1, 2, 3},
		{6, -3, 3},
		{0, 0, 0},
		{35, 35, 70},
		{-567, 500, -67},
		{0.05, 0.05, 0.1},
		{545.99, 0.01, 546},
		{100500, -100000, 500},
	}

	for _, tc := range tests {
		got := sum(tc.A, tc.B)
		if got != tc.want {
			t.Errorf("sum(%v, %v) = %v, want %v", tc.A, tc.B, got, tc.want)
		}
	}
}

func TestDivide(t *testing.T) {
	type test struct {
		A    float64
		B    float64
		want float64
	}

	tests := []test{
		{5, 1, 5},
		{25, -1, -25},
		{-9, -3, 3},
		{0, 5, 0},
		{-10, 2, -5},
		{1, 2, 0.5},
		{0.01, 0.01, 1},
		{30000000000, 1000, 30000000},
	}

	for _, tc := range tests {
		got := divide(tc.A, tc.B)
		if got != tc.want {
			t.Errorf("divide(%v, %v) = %v, want %v", tc.A, tc.B, got, tc.want)
		}
	}
}

func TestAverage(t *testing.T) {
	type test struct {
		A    float64
		B    float64
		want float64
	}

	tests := []test{
		{1, 2, 1.5},
		{0, -4, -2},
		{-100, -6, -53},
		{0, 1000, 500},
		{-10, 2.6, -3.7},
		{1.45, 2, 1.725},
		{0.001, 0.1, 0.0505},
		{0, 1000000, 500000},
	}

	for _, tc := range tests {
		got := average(tc.A, tc.B)
		if got != tc.want {
			t.Errorf("average(%v, %v) = %v, want %v", tc.A, tc.B, got, tc.want)
		}
	}
}
