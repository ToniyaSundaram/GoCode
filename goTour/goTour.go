package main

import (
	"fmt"
	"math"
)

func Pow(m, n, limit float64) float64 {
	if v := math.Pow(m, n); v < limit {
		return v
	} else {
		fmt.Printf("%v >= %v \n", v, limit)
	}
	return limit
}

func Sqrt(x float64) float64 {

	if x > 0 {
		z := float64(x / 2)
		temp := 0.0

		for i := 1; i <= 10; i++ {
			temp = z - (z*z-x)/(2*z)
			fmt.Println("z", z)
			if (temp - z) < 0 {
				if -(temp - z) < 0.000000001 {
					return temp
					break
				}
			}
			z = temp
		}
		return temp
	}
	return 0
}
func main() {
	// fmt.Println(
	// 	Pow(3, 3, 9),
	// 	Pow(2, 2, 7),
	// )

	///=======If and else ======
	n := float64(2)
	fmt.Println("Square root", Sqrt(n))
	fmt.Println("math root", math.Sqrt(n))

}
