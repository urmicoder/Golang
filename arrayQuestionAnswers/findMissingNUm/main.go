package main

import "fmt"

func findMissingElement(arr []int) int {
	if len(arr) == 0 {
		return -1
	}
	sumArr1 := 0
	n := len(arr) + 1
	//totalSum := n * (n + 1) / 2  yis ek line se vi same response milega
	totalSum := 0
	count := 0
	for _, val := range arr {
		sumArr1 += val
	}
	for i := 0; i < n; i++ {
		count++
		totalSum += count
	}
	return totalSum - sumArr1
}

func main() {
	// 	Input: [1, 2, 4, 5, 6]
	// Output: 3
	arrInt := []int{1, 2, 4, 5, 6}
	output := findMissingElement(arrInt)
	fmt.Println(output)
}
