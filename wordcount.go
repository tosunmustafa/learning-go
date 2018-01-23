package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	stripped := strings.Fields(s)
	results  := make(map[string]int)
	for _, s := range stripped {
		_ , ok := results[s]
		if(ok) {
			results[s] = results[s] + 1
		} else {
			results[s] = 1
		}
	}
	return results
}

func main() {
	wc.Test(WordCount)
}
