package main

import "fmt"

func findMaximumElement(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	maxVal := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
		}
	}
	return maxVal
}

func main() {
	//	Input: [3, 7, 2, 9, 5]
	// Output: 9
	arrInt := []int{3, 7, 2, 9, 5}
	output := findMaximumElement(arrInt)
	fmt.Println(output)
}
