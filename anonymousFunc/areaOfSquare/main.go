package main

import "fmt"

func main() {
	var area = func(side int) int {
		return side * side
	}
	result := area(4)
	fmt.Println("The area of square is", result)
}
