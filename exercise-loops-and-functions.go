package main

import (
	"fmt"
)

func Sqrt(x float64) float64 {
	/* z := float64(1) */
	z := 1.0
	for i := 1; i < 6; i++ {
		z = (z + x/z) / 2
	}

	return z
}

func main() {
	fmt.Println(Sqrt(2))
}
