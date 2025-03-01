package main

import "fmt"

func main() {
	values1 := []int{13, 52, 75, 96}
	values2 := []int{43, 95, 86, 32, 76}
	num1 := copy(values2, values1) //[13, 52, 75, 96,76]
	fmt.Println(values2, num1)     //[13 52 75 96 76] 4
}

/*

func main() {
	values1 := []int{13, 52, 75, 96, 24, 4}
	values2 := []int{43, 95, 86, 32, 76}
	num1 := copy(values2, values1) //[13, 52, 75, 96,76]
	fmt.Println(values2, num1)     //[13 52 75 96 24] 5   jitna values2 ki length hai utna hi element add hoga
}

func main() {
	values1 := []int{13, 52, 75, 96}
	values2 := []int{43, 95, 86, 32, 76}
	num1 := copy(values2, values1) //[13, 52, 75, 96,76]
	fmt.Println(values2, num1)     //[13 52 75 96 76] 4
}

func main() {
	src := []int{1, 2, 3, 4, 5}   MSys me Pucha gya tha
	dst := make([]int, 3) //[0,0,0]

	copied := copy(dst, src) //[1, 2, 3, 4, 5] wrong output
	fmt.Println(dst, copied) //[1 2 3] 3
}*/

// func main() {
// 	src := []int{1, 2, 3, 4, 5}
// 	dst := make([]int, 6) //[0,0,0,0,0,0]

// 	copied := copy(dst, src)
// 	fmt.Println(dst, copied) //[1 2 3 4 5 0] 5
// }
