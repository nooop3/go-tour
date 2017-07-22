package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return x, ErrNegativeSqrt(x)
	}
	z := float64(1)
	for math.Abs(x-z*z) >= 0.000001 {
		z = z - (z*z-x)/2/z
	}
	return z, nil
}

func main() {
	fmt.Println(math.Sqrt2)
	fmt.Println(Sqrt(-2))
	fmt.Println(Sqrt(2))
}
