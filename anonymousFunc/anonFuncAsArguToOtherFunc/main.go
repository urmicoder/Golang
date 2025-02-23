package main

import "fmt"

/* Anonymous Function as Arguments to Other Functions
Anonymous function ko ek argument ke roop me bhi kisi doosre function ko pass kiya ja sakta hai. */

func areaOfRectangle(l, w int) int {
	return l * w
}

func main() {
	sub := func(num1, num2 int) int {
		return num1 - num2
	}

	result := areaOfRectangle(sub(7, 3), sub(7, 4))
	fmt.Println("Area of Rectangle is", result)
}
