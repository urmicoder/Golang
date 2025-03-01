package main

import "fmt"

func firstNonZeroElement(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	for _, val := range arr {
		if val != 0 {
			return val
		}
	}
	return -1
}

func main() {
	//	Input: [0, 0, 5, 2, 0]
	//
	// Output: 5
	arrInt := []int{0, 0, 5, 2, 0}
	output := firstNonZeroElement(arrInt)
	fmt.Println(output)
}
