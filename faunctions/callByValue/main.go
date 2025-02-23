package main

import "fmt"

/* In Go, there are two ways to pass arguments to a function:

Pass by Value (Call by Value) – A copy of the value is passed to the function.
Pass by Reference (Call by Reference) – A reference (memory address) is passed instead of a copy.
By default, Go uses Call by Value.

In this approach, the actual values are copied inside the function.
Both the original and copied values are stored in different memory locations.

So, if the value changes inside the function, it does not affect the original value. */

func change(x int) {
	x = 100
}

func main() {
	x := 13
	fmt.Printf("Before the function call, value of x is = %d\n", x)
	change(x)
	fmt.Printf("After the function call, value of x is = %d\n", x)
}

// pc@hp:~/P/Golang/faunctions/callByValue$ go run main.go
// Before the function call, value of x is = 13
// After the function call, value of x is = 13
