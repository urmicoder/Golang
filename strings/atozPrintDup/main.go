package main

import "fmt"

func dupPrint(str string) {
	frequencyMap := make(map[rune]int)
	for _, char := range str {
		frequencyMap[char]++
	}
	for char, count := range frequencyMap {
		result := string(char)
		if count > 1 {
			fmt.Println(result, count)
		}
	}
}

func main() {
	str := "wqewrtri"
	dupPrint(str)
}
