package main

import "fmt"

/* Return an Anonymous Function from a Function
Anonymous function ko ek normal function se return bhi kiya ja sakta hai. */

func displayNumber() func() int {
	num := 2
	return func() int {
		num++
		return num
	}
}

func main() {
	a := displayNumber()

	fmt.Println(a())
}
