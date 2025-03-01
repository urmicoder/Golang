package main

import (
	"fmt"
	"strings"
	"unicode"
)

func capitalize(str string) string {
	words := strings.Fields(str)
	var result []string
	for _, word := range words {
		n := len(word)
		if n > 1 {
			first := string(unicode.ToUpper(rune(word[0])))
			last := string(unicode.ToUpper(rune(word[n-1])))
			middle := (word[1:(n - 1)])
			result = append(result, first+middle+last)
		} else {
			result = append(result, strings.ToUpper(word))
		}
	}
	return strings.Join(result, " ")
}

func main() {
	str := "take u forward is awesome"
	outputStr := capitalize(str)
	fmt.Printf("%q \n", outputStr)
}
