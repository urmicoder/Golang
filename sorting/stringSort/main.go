package main

import (
	"fmt"
	"sort"
)

// acending order
func main() {
	sliceString := []string{"urmi", "durga", "kewat", "coder"}
	fmt.Println("Slice of string Before sort", sliceString)
	sort.Strings(sliceString)
	fmt.Println("Slice of string After sort", sliceString)
}

// pc@hp:~/P/Golang/sorting/stringSort$ go run main.go
// Slice of string Before sort [urmi durga kewat coder]
// Slice of string After sort [coder durga kewat urmi]
