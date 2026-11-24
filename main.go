package main

import (
	"fmt"
	"strings"
)

// ReverseWords reverses the order of the words in s, collapsing any runs of
// whitespace to single spaces and trimming the ends.
func ReverseWords(s string) string {
	words := strings.Fields(s)
	for l, r := 0, len(words)-1; l < r; l, r = l+1, r-1 {
		words[l], words[r] = words[r], words[l]
	}
	return strings.Join(words, " ")
}

func main() {
	fmt.Printf("%q\n", ReverseWords("  the sky   is blue "))
}
