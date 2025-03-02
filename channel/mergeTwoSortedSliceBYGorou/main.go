package main

import "fmt"

func marged(left, right []int, ch chan []int) {
	result := []int{}
	i, j := 0, 0
	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	ch <- result
}

func main() {
	left := []int{1, 3, 5}
	right := []int{2, 4, 6}
	ch := make(chan []int)
	go marged(left, right, ch)
	data := <-ch
	fmt.Println(data) //1 2 3 4 5 6]
}
