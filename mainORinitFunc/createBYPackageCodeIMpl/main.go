package main

/* func fun() int {
	println("this is func")
	return 1
}

func init() {
	println("Inside main.go init function.")
}

func main() {
	println("Inside the main function.")
	println("Value of x is ", fun())
}

Packages load hote hain
2️⃣ Global variables initialize hote hain
3️⃣ init() function execute hota hai
4️⃣ Finally, main() function execute hota hai

pc@hp:~/P/Golang/mainORinitFunc/createBYPackageCodeIMpl$ go run main.go
Inside main.go init function.
Inside the main function.
this is func
Value of x is  1  */

var x = fun()

func fun() int {
	println("this is func")
	return 1
}

func init() {
	println("Inside main.go init function.")
}

func main() {
	println("Inside the main function.")
	println("Value of x is ", x)
}

// pc@hp:~/P/Golang/mainORinitFunc/createBYPackageCodeIMpl$ go run main.go
// this is func
// Inside main.go init function.
// Inside the main function.
// Value of x is  1
