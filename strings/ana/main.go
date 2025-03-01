package main

import (
	"fmt"
	"sort"
)

func anagram(str1, str2 string) bool {
	runes1 := []rune(str1) //[65 67 84]
	runes2 := []rune(str2) //[67 65 84]
	sort.Slice(runes1, func(i, j int) bool {
		return runes1[i] > runes1[j]
	})
	sort.Slice(runes2, func(i, j int) bool {
		return runes2[i] > runes2[j]
	})
	return string(runes1) == string(runes2)
}

func main() {
	str1 := "ACT"
	str2 := "CAT"
	output := anagram(str1, str2)
	fmt.Println("string is anagram:", output)
}
