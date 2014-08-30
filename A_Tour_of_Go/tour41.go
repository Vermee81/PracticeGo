package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	ans := make(map[string]int)
	array := strings.Fields(s)
	for _, word := range strings.Fields(s) {
		ans[word] += 1
	}
	return ans
}

func main() {
	wc.Test(WordCount)
}
