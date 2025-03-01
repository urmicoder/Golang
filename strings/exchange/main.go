package main

import "fmt"

func exchangeValue(str1, str2 string) (string, string) {
	return str2, str1
}

func main() {
	str1 := "urmi"
	str2 := "rupam"
	output1, output2 := exchangeValue(str1, str2)
	fmt.Println("Swap string Value without temp variable:", output1, output2)
}
