package main

import "fmt"

//we can extract some portion from current slice using [lo:hi] where lo start from index and go upto to h1-1.

func main() {
	myslice := []int{3, 5, 8, 12, 9, 11, 45, 32}
	fmt.Println(myslice)
	fmt.Println("MySlice[2:4]=", myslice[2:4]) //[8,12]
	fmt.Println("myslice[:5]=", myslice[:5])   //[3, 5, 8, 12, 9]
	fmt.Println("myslice[4:8]=", myslice[4:8]) //[9, 11, 45, 32]
	fmt.Println("myslice2=", myslice[3:])      //[12, 9, 11, 45, 32]
}
