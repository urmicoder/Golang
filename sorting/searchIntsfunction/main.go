package main

import (
	"fmt"
	"sort"
)

/* SearchInts function [ascending order]
SearchInts function ek sorted integer slice mein x ki position ko search karta hai aur index return karta hai.
 Yeh function tabhi kaam karta hai jab slice pehle se sorted ho. */

func main() {
	intSlice := []int{45, 12, 28, 19, 14, 88, 28, 16, 55, 65}
	x := 16
	pos := sort.SearchInts(intSlice, x)
	fmt.Println(pos)
	sort.Ints(intSlice)
	pos = sort.SearchInts(intSlice, x)
	fmt.Printf("Present %d at index %d in %v\n", x, pos, intSlice)
}
