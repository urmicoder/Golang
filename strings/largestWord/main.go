package main

import (
	"fmt"
	"strings"
)

func findLargestWord(str string) string {
	words := strings.Fields(str)
	result := " "
	maxLengthCount := 0
	for _, word := range words {
		maxlength := len(word)
		if maxlength > maxLengthCount {
			maxLengthCount = maxlength
			result = word
		}
	}
	return result
}

func main() {
	str := "Google Doc" //output = "Google"
	output := findLargestWord(str)
	fmt.Println(output)
}
