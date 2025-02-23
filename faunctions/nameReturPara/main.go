package main

import "fmt"

/*  Named return variables are defined inside round brackets () after the function parameters.

Here, a, b, and c are already declared, so we don’t need to specify them in the return statement! */

func exampalName(A, B, C int) (a, b, c int) {
	a = A + 5
	b = B + 10
	c = C + 15
	return
}

func main() {
	x, y, z := exampalName(5, 10, 15)
	fmt.Println(x, y, z)
}
