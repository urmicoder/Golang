package main

import "fmt"

func main() {
	var char rune
	fmt.Println("Enter Somthing")
	fmt.Scanf("%c", &char)
	value := int(char)
	fmt.Printf("Value is %c %d\n", char, value)
}
