package main

import (
	"fmt"
	"strings"
)

func countVCS(str string) (SpaceCount, vowelsCount, capitalCount int) {
	vowels := "aeiou"

	for _, char := range str {
		if char == ' ' {
			SpaceCount++
		} else if strings.ContainsRune(vowels, char) {
			vowelsCount++
		} else if char >= 'a' && char <= 'z' {
			capitalCount++
		}
	}
	return
}

func main() {
	str := "ksh fiu regry"
	SpaceCount, vowelsCount, capitalCount := countVCS(str)
	fmt.Println(SpaceCount, vowelsCount, capitalCount)
}
