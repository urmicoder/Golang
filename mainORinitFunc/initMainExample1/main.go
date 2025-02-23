package main

import "fmt"

/*  Yeh main function se pehle execute hota hai.
✅ Ek program me multiple init() functions likhne ki permission hoti hai.
✅ Yeh lexical order me execute hota hai.
✅ Is function ko manually call nahi kar sakte, agar kiya toh error aayega.
............Order of Execution
Golang me execution ek particular order me hoti hai:

1️⃣ Saare required packages ko import kiya jata hai.
2️⃣ Imported packages ko initialize kiya jata hai.
3️⃣ Constant variables initialize hote hain.
4️⃣ Baaki saare variables initialize hote hain.
5️⃣ Sabse last me init() functions execute hote hain. */ //matlob main se nhi call kar shakte

func init() {
	fmt.Println("Hello init function")
}

func main() {
	fmt.Println("Welcome to new world")
}

// pc@hp:~/P/Golang/mainORinitFunc/initMainExample1$ go run main.go
// Hello init function
// Welcome to new world
