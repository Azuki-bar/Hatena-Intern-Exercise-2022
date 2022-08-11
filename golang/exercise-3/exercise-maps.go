package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := make(map[string]int)
	for _, word := range words {
		count[word] += 1
		// if c, ok := count[word]; ok {
		// 	count[word] = c + 1
		// } else {
		// 	count[word] = 1
		// }
	}
	return count
}

func main() {
	wc.Test(WordCount)
}
