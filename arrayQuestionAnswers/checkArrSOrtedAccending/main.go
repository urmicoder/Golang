package main

import (
	"fmt"
)

func checkArraySortedAsscending(arrInt []int) bool {
	//sort.Ints(arrInt)  yis ek line se vi same response milega

	if len(arrInt) == 0 {
		return false
	}
	for i := 1; i < len(arrInt); i++ {
		if arrInt[i] < arrInt[i-1] {
			return false
		}
	}
	return true
}

func main() {
	// 	Input: [1, 2, 3, 4]
	// Output: true
	arrInt := []int{1, 2, 3, 4}
	output := checkArraySortedAsscending(arrInt)
	fmt.Println(output)
}
