package main

import "fmt"

//func can retrun multipal values

func swapData(str1, str2 string) (string, string) {
	return str2, str1
}

func main() {
	str1, str2 := "urmi", "rupam"
	x, y := swapData(str1, str2)
	fmt.Println(x, y)
}
