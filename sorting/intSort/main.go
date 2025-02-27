package main

import (
	"fmt"
	"sort"
)

/* Sorting means arranging elements in an ordered sequence.

Slice sorting means working with slices and organizing them in different ways. In Golang,
we can sort any list using built-in sort functions. To use sorting, we need to import the "sort" package at the beginning of our code. */

func main() {
	sliceInt := []int{10, 5, 25, 646, 14, 4}
	fmt.Println("Slice of integer Before sort", sliceInt)
	sort.Ints(sliceInt)
	fmt.Println("Slice of integer After sort", sliceInt)
}

// pc@hp:~/P/Golang/sorting/intSort$ go run main.go
// Slice of integer Before sort [10 5 25 646 14 4]
// Slice of integer After sort [4 5 10 14 25 646]
