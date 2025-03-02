package main

import (
	"fmt"
	"strings"
)

func hightNum(str string) string {
	words := strings.Fields(str)
	result := " "
	maxReapted := 0
	maxCharCount := 0
	for _, word := range words {
		frequencyMap := make(map[rune]int)
		for _, char := range word {
			frequencyMap[char]++
			if frequencyMap[char] > maxCharCount {
				maxCharCount = frequencyMap[char]
			}
		}
		if maxCharCount > 1 && maxCharCount > maxReapted {
			maxReapted = maxCharCount
			result = word
		}
	}
	if result == " " {
		return "-1"
	}
	return result
}

func main() {
	str := "abcdefghi google microsoft" //output = "google"
	//fmt.Println(str)
	output := hightNum(str)
	fmt.Println(output)

}
