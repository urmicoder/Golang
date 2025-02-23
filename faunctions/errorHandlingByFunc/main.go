package main

import (
	"errors"
	"fmt"
)

/* Multiple return values make error handling easier because we can check errors one by one.
This keeps the program flow simple and clear. */

func fun(a int) (int, error) {
	if a == 0 {
		return 0, errors.New("zeros are not allowed")
	} else {
		return a * 2, nil
	}
}

func main() {
	x, err := fun(4)
	if err != nil {
		fmt.Println("Error Detected")
	} else {
		fmt.Println(x)
	}
}
