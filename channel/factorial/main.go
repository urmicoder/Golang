package main

import "fmt"

func maxNumber(arr []int, maxCh chan int) {
	maxNum := arr[0]
	for i := 0; i < len(arr); i++ {
		if arr[i] > maxNum {
			maxNum = arr[i]
		}
	}
	maxCh <- maxNum
	//return maxNum
}

func factorial(maxNum int, factorialCh chan int) {
	result := 1
	for i := 2; i < maxNum; i++ {
		result *= i
	}
	factorialCh <- result
}

func main() {
	arr := []int{2, 5, 6, 7, 12, 9}
	maxCh := make(chan int)
	factorialCh := make(chan int)

	go maxNumber(arr, maxCh)
	data := <-maxCh
	go factorial(data, factorialCh)

	data2 := <-factorialCh
	fmt.Println(data2)
}
