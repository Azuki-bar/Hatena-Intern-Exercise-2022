package main

import (
	"fmt"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %.0f", float64(e))
}
func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	z := 1.0
	count := 1
	allowedRange := 0.00001
	for {
		dz := (z*z - x) / (2 * z)
		z -= dz
		if -allowedRange <= dz && dz <= allowedRange {
			return z, nil
		}
		count++
		if count%10 == 0 {
			fmt.Println(z)
		}
	}
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
