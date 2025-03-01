package main

import "fmt"

//A slice is like an array, but it is more powerful, flexible, and easy to use. An array has a fixed size,
//but a slice can grow or shrink as needed. A slice is a flexible way to view and use elements of an array.

func main() {
	var sliceInt []int                                                  //Declare a Slice Without Initialization
	fmt.Printf("Length=%d Capacity=%d\n", len(sliceInt), cap(sliceInt)) //Length=0 Capacity=0
	fmt.Println(sliceInt == nil)
}

// func main() {
// 	slicesInt := []int(nil)
// 	fmt.Println(slicesInt == nil)
// }

// pc@hp:~/P/Golang/slices/nilSlice$ go run main.go
// true
