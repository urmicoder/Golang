package main

import (
	"fmt"
	"strings"
)

func occurenceIndexValue(str1, str2 string) int {
	//words := strings.Fields(str1)
	data := strings.Index(str1, str2)
	return data
}

func main() {
	str1 := "takeurforward" //basically yisme us word ka occrrencs find karna hai jo dono string me present ho.
	str2 := "forward"
	outputInt := occurenceIndexValue(str1, str2)
	fmt.Println(outputInt) //6
}
