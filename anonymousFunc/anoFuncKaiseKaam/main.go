package main

import "fmt"

/*Agar function ka naam nahi hai, toh use call kaise karenge?

Simple! Anonymous function ko ek variable me assign kar dete hain, aur phir us variable se function ko call karte hain.

func main() {
	func() {
		fmt.Println("Hello, Welcome to interview place")
	}()
} */

func main() {
	var welcome = func() {
		fmt.Println("Hello, Welcome to interview place")
	}

	welcome()
}

// pc@hp:~/P/Golang/anonymousFunc/anoFuncKaiseKaam$ go run main.go
// Hello, Welcome to interview place
// pc@hp:~/P/Golang/anonymousFunc/anoFuncKaiseKaam$ go run main.go
// Hello, Welcome to interview place
// pc@hp:~/P/Golang/anonymousFunc/anoFuncKaiseKaam$ go run main.go
// Hello, Welcome to interview place
