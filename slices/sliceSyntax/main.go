package main

import "fmt"

func main() {
	sliceInt := make([]int, 3)
	sliceInt[0] = 2
	sliceInt[1] = 3
	sliceInt[2] = 4
	fmt.Printf("Values of sliceint %d\n", sliceInt) //Values of sliceint [2 3 4]
	sliceInt2 := make([]int, 5, 8)
	fmt.Printf("Lenght = %d\n", len(sliceInt2))
	fmt.Printf("Capacity = %d\n", cap(sliceInt2))
	fmt.Printf("Value = %d\n", sliceInt2)
}

// pc@hp:~/P/Golang/slices/sliceSyntax$ go run main.go
// Values of sliceint [2 3 4]
// Lenght = 5
// Capacity = 8
// Value = [0 0 0 0 0]
