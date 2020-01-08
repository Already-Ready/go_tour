package main

import (
	"strings"

	"code.google.com/p/go-tour/wc"
)

func WordCount(s string) map[string]int {
	m := make(map[string]int)
	for _, v := range strings.Fields(s) {
		_, b := m[v]
		if b == true {
			m[v] += 1
		} else if b == false {
			m[v] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}
