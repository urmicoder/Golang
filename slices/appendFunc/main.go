package main

import "fmt"

//The append() function allows us to add new elements at the end of a slice.

// func main() {
// 	values := make([]int, 4)                                                         //[0,0,0,0]
// 	values = append(values, 5)                                                       //[0,0,0,0,5]
// 	values = append(values, 9, 12, 56, 2)                                            //[0,0,0,0,5, 9, 12, 56, 2]
// 	fmt.Printf("Length=%d Capacity=%d Value=%v\n", len(values), cap(values), values) //Length=9 Capacity=16 Value=[0 0 0 0 5 9 12 56 2]
// 	values1 := []int{10, 12, 25}
// 	values2 := []int{45, 20, 55, 65}
// 	values3 := append(values1, values2...)
// 	fmt.Printf("Length=%d Capacity=%d Values=%v\n", len(values), cap(values), values3) //Length=7 Capacity=12 Value=[10, 12, 25, 45, 20, 55, 65]
// }

func main() {
	values1 := []int{10, 12, 25, 11, 10}
	values2 := []int{45, 20, 55, 65}
	values3 := append(values1, values2...)
	fmt.Printf("Length=%d Capacity=%d Values=%v\n", len(values3), cap(values3), values3) //Length=9 Capacity=10 Values=[10 12 25 11 10 45 20 55 65]
}

// func main() {
// 	values1 := []int{10, 12, 25}
// 	values2 := []int{45, 20, 55, 65}
// 	values3 := append(values1, values2...)
// 	fmt.Printf("Length=%d Capacity=%d Values=%v\n", len(values3), cap(values3), values3) //Length=7 Capacity=8 Values=[10 12 25 45 20 55 65]
// }

/*


func main() {
	values1 := []int{10, 12}
	values2 := []int{45, 20, 55, 65}
	values3 := append(values1, values2...)
	fmt.Printf("Length=%d Capacity=%d Values=%v\n", len(values3), cap(values3), values3) //Length=6 Capacity=6 Values=[10 12 45 20 55 65]
}

func main() {
	values1 := []int{}
	values2 := []int{45, 20, 55, 65}
	values3 := append(values1, values2...)
	fmt.Printf("Length=%d Capacity=%d Values=%v\n", len(values3), cap(values3), values3) //Length=4 Capacity=4 Values=[45 20 55 65]
} */

//cap ko even number rakhana hai
