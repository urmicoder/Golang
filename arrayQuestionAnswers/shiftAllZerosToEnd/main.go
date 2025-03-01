package main

import "fmt"

func shiftZeroToEnd(arr []int) []int {
	j := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			j++
		}
	}
	return arr
}

func main() {
	//	Input: [1, 0, 2, 0, 3]
	// Output: [1, 2, 3, 0, 0]
	arrInt := []int{1, 0, 2, 0, 3}
	output := shiftZeroToEnd(arrInt)
	fmt.Println(output)
}
