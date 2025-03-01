package main

import "fmt"

func countOccurrences(arr []int, target int) int {
	if len(arr) == 0 {
		return -1
	}
	var count = 0
	for _, val := range arr {
		if val == target {
			count++
		}
	}
	return count
}

func main() {
	// 	Input: [1, 2, 3, 1, 2, 1, 4], Find 1
	// Output: 3
	arrInt := []int{1, 2, 3, 1, 2, 1, 4}
	target := 1
	output := countOccurrences(arrInt, target)
	fmt.Println(output)

}
