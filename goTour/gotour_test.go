package main

import (
	"math"
	"testing"
)

func TestGo(t *testing.T) {
	// tables := []struct {
	// 	m     float64
	// 	n     float64
	// 	limit float64
	// }{
	// 	{3, 3, 20},
	// 	{5, 5, 453},
	// 	{2, 2, 4},
	// 	{7, 7, 1234},
	// }

	// for _, table := range tables {
	// 	result := Pow(table.m, table.n, table.limit)
	// 	fmt.Println(result)
	// 	if result != table.limit {
	// 		t.Errorf("The result should be limit %v", table.limit)
	// 	}
	// }

	var cases [5]float64
	//positive
	cases[0] = 1024
	// zero
	cases[2] = 0

	for _, c := range cases {
		result := Sqrt(c)
		if result != math.Sqrt(c) {
			t.Errorf("The calculated square root is not correct")
		}
	}

}
