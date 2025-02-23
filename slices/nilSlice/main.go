package main

import "fmt"

// func main() {
// 	var sliceInt []int   Declare a Slice Without Initialization
// 	fmt.Println(sliceInt == nil)
// }

func main() {
	slicesInt := []int(nil)
	fmt.Println(slicesInt == nil)
}

// pc@hp:~/P/Golang/slices/nilSlice$ go run main.go
// true
