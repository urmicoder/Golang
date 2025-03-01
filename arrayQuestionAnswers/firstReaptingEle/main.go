package main

import "fmt"

func firstRepeatingElement(arr []int) int {
	freqMap := make(map[int]bool)
	for _, val := range arr {
		if freqMap[val] {
			return val
		}
		freqMap[val] = true
	}
	return -1
}

func main() {
	// 	Input: [4, 5, 2, 4, 3, 5]
	// Output: 4
	arrInt := []int{4, 5, 2, 4, 3, 5}
	output := firstRepeatingElement(arrInt)
	fmt.Println(output)
}
