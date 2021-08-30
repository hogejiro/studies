package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	ret := make(map[string]int)
	words := strings.Fields(s)
	for _, word := range words {
		ret[word] = ret[word] + 1
	}
	return ret
}

func main() {
	wc.Test(WordCount)
}
