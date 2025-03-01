package main

import "fmt"

func sumOfAllElement(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

func main() {
	//	Input: [4, 2, 1, 6]
	// Output: 13
	arrInt := []int{4, 2, 1, 6}
	output := sumOfAllElement(arrInt)
	fmt.Println(output)
}
