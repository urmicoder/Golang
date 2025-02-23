package main

import "fmt"

// func main() {

// 	func(a, b int) int {
// 		sum := a + b
// 		fmt.Println(sum)        //with Retrun value
// 		return sum
// 	}(2, 3)
// }

func main() {

	var sum = func(a, b int) {
		sum := a + b
		fmt.Println("addtion is", sum) //without Retrun value
	}
	sum(2, 3)
}

// pc@hp:~/P/Golang/anonymousFunc/anoFuncWithPara$ go run main.go
// 5
// pc@hp:~/P/Golang/anonymousFunc/anoFuncWithPara$ go run main.go
// addtion is 5
