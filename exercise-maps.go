package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	re := make(map[string]int)
	sArr := strings.Fields(s)
	for _, w := range sArr {
		if _, ok := re[w]; ok {
			re[w] += 1
		} else {
			re[w] = 1
		}
	}
	return re
}

func main() {
	wc.Test(WordCount)
}
