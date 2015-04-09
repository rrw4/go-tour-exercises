package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	val := make(map[string]int)
	words := strings.Fields(s)
	for _, v := range words {
		elem, ok := val[v]
		if ok {
			val[v] = elem + 1
		} else {
			val[v] = 1
		}
	}
	return val
}

func main() {
	wc.Test(WordCount)
}
