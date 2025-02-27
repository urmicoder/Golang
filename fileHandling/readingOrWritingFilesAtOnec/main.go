package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

//Reading and Writing Files
//To read files from the local filesystem, we use the io/ioutil module in Go. This module helps us load some file content into memory.

func CreateFile() {
	file, err := os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	len, err := file.WriteString("Welcome to kolkata" + " " + "And Celibrate Durga Puja Here")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nFile Name: %s", file.Name())
	fmt.Printf("\nFile Length: %d", len)
}

func ReadFile() {
	fileName := "test.txt"

	data, err := ioutil.ReadFile("test.txt") //Reads the entire file at once, which may not be efficient for large files.
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("\nFile Name: %s", fileName)
	fmt.Printf("\nFile Length: %d", len(data))
	fmt.Printf("\nData: %s\n", data)
}

func main() {
	CreateFile()
	ReadFile()
}

// pc@hp:~/P/Golang/fileHandling/readingOrWritingFiles$ go run main.go

// File Name: test.txt
// File Length: 48
// File Name: test.txt
// File Length: 48
// Data: Welcome to kolkata And Celibrate Durga Puja Here
