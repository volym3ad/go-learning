package main

import (
	"fmt"
	"strings"
)

func WordCount(s string) map[string]int {
	words := strings.Fields(s)
	count := map[string]int{}
	for _, word := range words {
		count[word]++
	}
	return count
}

func main() {
	test := map[string]int{}
	test = WordCount("dadfds dfsdf dfgds sdf s ff sdf ff")
	fmt.Printf("%#v", test)
}