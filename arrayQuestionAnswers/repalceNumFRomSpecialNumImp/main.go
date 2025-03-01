package main

import "fmt"

func main() {
	//[4, 2, 4, 3], Replace 4 with 0

	arrInt := []int{4, 2, 4, 3}
	replace := 4
	for i, _ := range arrInt {
		if arrInt[i] == replace {
			arrInt[i] = 0
			//fmt.Println(arrInt)
		}
	}
	fmt.Println(arrInt) //arr me value menipulation index se hi kar shakte direct val se nhi
	//[0 2 0 3]
}
