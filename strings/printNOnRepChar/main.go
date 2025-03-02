package main

import "fmt"

func nonREpChar(str string) {
	frequencyMap := make(map[rune]int)

	for _, char := range str {
		frequencyMap[char]++
	}
	for char, count := range frequencyMap {
		// if count < 2 {
		// 	fmt.Printf("%q \n", string(char))
		// }
		if count == 1 {
			fmt.Printf("%c\n", char)
		}
	}
}

func main() {
	str := "google"
	nonREpChar(str)
	//fmt.Println(str)
}
