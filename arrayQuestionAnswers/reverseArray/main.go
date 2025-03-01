package main

import "fmt"

/*
func reverseArray(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	i := 0
	j := n
	for i < j {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
		i++
		j--
	}
	return arr
}*/

func reverseArray(arr []int) []int {
	n := len(arr)
	if n == 0 {
		return []int{}
	}
	for i := 0; i < n/2; i++ {
		arr[i], arr[n-1-i] = arr[n-1-i], arr[i]
	}
	return arr
}

func main() {
	//	input: [1, 2, 3, 4]
	// Output: [4, 3, 2, 1]
	arrInt := []int{1, 2, 3, 4}
	output := reverseArray(arrInt)
	fmt.Println(output)

}
