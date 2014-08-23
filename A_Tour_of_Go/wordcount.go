package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	splitted := strings.Fields(s)
	for _, str := range splitted {
		m[str] = m[str] + 1
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
