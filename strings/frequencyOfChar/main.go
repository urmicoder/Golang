package main

import "fmt"

func frequencyOfCharValue(str string) {
	frequencyMap := make(map[rune]int)

	for _, char := range str {
		frequencyMap[char]++
	}

	for char, count := range frequencyMap {
		result := string(char)
		fmt.Println(result, count)
	}
}

func main() {
	str := "urmirupam"
	frequencyOfCharValue(str)
}
