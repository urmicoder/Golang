package main

import (
	"fmt"
	"log"
	"os"
)

//Creating Multiple Directories
//MkdirAll function ek directory create karta hai, saath hi uske necessary parent directories bhi. Ye nil ya error message return karta hai.

func main() {
	path := "temp/golang/test"

	err := os.MkdirAll(path, 0755)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Path has been created")
}

// pc@hp:~/P/Golang/fileHandling/multipalDire$ go run main.go
// Path has been created
