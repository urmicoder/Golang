package main

import "fmt"

//A slice is like an array, but it is more powerful, flexible, and easy to use. An array has a fixed size,
//but a slice can grow or shrink as needed. A slice is a flexible way to view and use elements of an array.

func main() {
	sliceInt := make([]int, 0)
	fmt.Println(sliceInt)
	sliceInt2 := []int{} //Yahan sliceInt2 ek empty slice hai jiska length aur capacity 0 hai
	fmt.Println(sliceInt2)
}

// pc@hp:~/P/Golang/slices/emptySlice$ go run main.go
// []
// []
