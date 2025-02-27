package main

import (
	"fmt"
	"sort"
)

func main() {
	a := []int{17, 5, 34, 55, 531, 98, 7, 12, 15, 105}
	dat := sort.Reverse(sort.IntSlice(a))
	fmt.Println(dat) //&{[17 5 34 55 531 98 7 12 15 105]}
	sort.Sort(dat)
	fmt.Println(a) //[531 105 98 55 34 17 15 12 7 5]

	a = []int{-17, -5, -34, -55, -531, -98, -7, -12, -15, -105}
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	fmt.Println(a) //[-5 -7 -12 -15 -17 -34 -55 -98 -105 -531]
}
