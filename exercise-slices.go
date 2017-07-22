package main

import (
	// "fmt"
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	re := make([][]uint8, dy)
	for i := 0; i < dy; i++ {
		r := make([]uint8, dx)
		for j := 0; j < dx; j++ {
			r[j] = uint8(i * j)
		}
		re[i] = r
	}

	// fmt.Println(dy)
	// fmt.Println(dx)
	return re
}

func main() {
	// fmt.Println(Pic(8, 8)
	pic.Show(Pic)
}
