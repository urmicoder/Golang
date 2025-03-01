package main

import "fmt"

func secondLargestElement(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	maxInt := arr[0]
	secoundMaxInt := arr[0]

	for i := 0; i < len(arr); i++ {
		if arr[i] > maxInt {
			secoundMaxInt = maxInt
			maxInt = arr[i]
		} else {
			if arr[i] > secoundMaxInt && arr[i] != maxInt {
				secoundMaxInt = arr[i]
			}
		}
	}
	return secoundMaxInt
}

func main() {
	// 	Input: [5, 3, 9, 7, 9]
	// Output: 7
	arrInt := []int{5, 3, 9, 7, 9}
	output := secondLargestElement(arrInt)
	fmt.Println(output)

}
