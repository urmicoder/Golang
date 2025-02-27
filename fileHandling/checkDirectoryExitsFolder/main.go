package main

import (
	"fmt"
	"log"
	"os"
)

// Check if the Directory exists
/*func main() {
	path := "temp"

	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0755)
		fmt.Println("Directory created")
	} else {
		fmt.Println("Directory already exists")
	}
}*/

// pc@hp:~/P/Golang/fileHandling/checkDirectoryExits$ go run main.go
// Directory created

//Deleting a Directory

func main() {
	path := "temp"

	err := os.RemoveAll(path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Remove Successful")
}

// pc@hp:~/P/Golang/fileHandling/checkDirectoryExits$ go run main.go
// Remove Successful
