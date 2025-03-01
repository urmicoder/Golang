package main

import (
	"fmt"
	"unicode"
)

func changeCase(str string) string {
	result := " "
	for _, char := range str {
		if char >= 'a' && char <= 'z' {
			result += string(unicode.ToUpper(char))
		} else if char >= 'A' && char <= 'Z' {
			result += string(unicode.ToLower(char))
		}
	}
	return result
}

func main() {
	str := "javA" //output = JAVa
	outputStr := changeCase(str)
	fmt.Println(outputStr)
}
