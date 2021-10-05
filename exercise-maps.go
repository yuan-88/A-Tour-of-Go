package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	sSplit := strings.Fields(s)
	for _, char := range sSplit {
		m[char]++
	}
	return m
}

func main() {
	wc.Test(WordCount)
}