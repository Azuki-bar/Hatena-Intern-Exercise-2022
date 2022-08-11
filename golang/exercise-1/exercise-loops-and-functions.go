package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	z := 1.0
	count := 1
	allowedRange := 0.00001
	for {
		dz := (z*z - x) / (2 * z)
		z -= dz
		if -allowedRange <= dz && dz <= allowedRange {
			return z
		}
		count++
		if count%10 == 0 {
			fmt.Println(z)
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
}
