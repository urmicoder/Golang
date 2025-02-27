package main

import (
	"fmt"
	"log"
	"os"
)

/*What is a Directory?
A directory is a unit in a computer’s file system that helps store and organize files. Directories are arranged in a hierarchical tree structure, where they follow a parent-child relationship. Directories are also called folders.
A directory is just a folder in Windows where we store our files. We can create, delete, and rename them as needed.

 Creating a Directory
 In Go, we use os.Mkdir to create a new directory with a specified name and permissions. */

func main() {
	err := os.Mkdir("golang_test", 0755) //Code 0755 ka matlab hai ki sabko read aur execute access milta hai, aur user ko write access milta hai.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("folder has been created")
}

// pc@hp:~/P/Golang/fileHandling/directory$ go run main.go
// folder has been created
