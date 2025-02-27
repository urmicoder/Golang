package main

import (
	"fmt"
	"sort"
)

func main() {
	sliceString := []string{"Coding", "Blogs", "Ninjas", "Vertical", "Data", "Structure", "New Delhi"}

	sort.Sort(sort.Reverse(sort.StringSlice(sliceString)))
	fmt.Println(sliceString) //[Vertical Structure Ninjas New Delhi Data Coding Blogs]
	c := []float64{70.10, 50.10, 140.15, 80.15, 6.95}
	sort.Sort(sort.Reverse(sort.Float64Slice(c)))
	fmt.Println(c) //[140.15 80.15 70.1 50.1 6.95]
}
