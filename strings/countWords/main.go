package main

import (
	"fmt"
	"strings"
)

func countOfWord(str string) int {
	words := strings.Fields(str)

	n := len(words)
	return n
}

func main() {
	str := "Hello, how are you doing today riya"
	outputStr := countOfWord(str)
	fmt.Println(outputStr)
}
