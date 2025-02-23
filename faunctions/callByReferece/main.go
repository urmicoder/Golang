package main

import "fmt"

/* Here, we pass the address of the variable instead of just the value. This uses pointers.

Since both the original and function parameters point to the same memory location,
any changes inside the function also affect the original value. */

func change(x *int) {
	*x = 100
}

func main() {
	var x = 13
	fmt.Printf("Before the Function Call, value of X is = %d\n", x)
	change(&x)
	fmt.Printf("After the Function Call, value of X is = %d\n", x)
}

// pc@hp:~/P/Golang/faunctions/callByReferece$ go run main.go
// Before the Function Call, value of X is = 13
// After the Function Call, value of X is = 100
