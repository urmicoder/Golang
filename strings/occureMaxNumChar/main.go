package main

import "fmt"

func occurMaxNum(str string) string {
	frequencyMap := make(map[rune]int)
	maxCountChar := 0
	var maxChar rune

	for _, char := range str {
		frequencyMap[char]++
		if frequencyMap[char] > maxCountChar {
			maxCountChar = frequencyMap[char]
			maxChar = char
		}
	}
	return string(maxChar)
}

func main() {
	str := "urmirupamm" //output = m
	output := occurMaxNum(str)
	fmt.Println(output)
}
