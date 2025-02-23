package main

import "fmt"

func main() {

	var sum = func(a, b int) int {
		return a + b
	}
	result := sum(2, 3)
	fmt.Println("addtion is", result)
}

// func main() {

// 	func(a, b int) int {
// 		sum := a + b
// 		fmt.Println(sum)        //with Retrun value
// 		return sum
// 	}(2, 3)
// }

// func main() {

// 	var sum = func(a, b int) int {
// 		sum := a + b
// 		fmt.Println("addtion is", sum) //with Retrun value
// 		return sum
// 	}
// 	sum(2, 3)
// }
