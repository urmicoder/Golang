package main

import "fmt"

/* A function is a small block of code that performs a specific task based on input data.
Using functions, we can repeat an operation without writing the same code again and again.

Input values in a function are optional, and the function may or may not return an output.
In Go, we use the func keyword to create a function. */

func definition() {
	fmt.Println("Welcome to function World")
}

func main() {
	definition()
}
