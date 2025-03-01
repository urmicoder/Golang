package main

import "fmt"

func modifiSlice(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	if n > 0 {
		arr[n-1] = 0
	}
	arr[n-1], arr[n-2] = arr[n-2], arr[n-1]
	return arr
}

func main() {
	inputArr := []int{1, 0, 2, 3}
	output := modifiSlice(inputArr)
	fmt.Println(output) //[1 0 0 2]
}
