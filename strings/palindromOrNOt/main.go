package main

import (
	"fmt"
	"strings"
)

func palindromString(str string) bool {
	str = strings.ReplaceAll(str, " ", "")
	n := len(str)
	for i := 0; i < n/2; i++ {
		if str[i] != str[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	str := "ABCD CBA"
	value := palindromString(str)
	fmt.Println(value)
}
