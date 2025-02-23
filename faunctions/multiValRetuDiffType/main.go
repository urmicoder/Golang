package main

import "fmt"

func defReturnType(a int) (int, float64) {
	return a / 2, float64(a) * 2.05
}

func main() {
	x, y := defReturnType(4)
	fmt.Println(x, y)
}
