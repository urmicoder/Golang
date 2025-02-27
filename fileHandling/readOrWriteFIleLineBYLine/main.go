package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func LineByLine(file string) error {

	fd, err := os.Open(file)

	if err != nil {
		return err
	}

	reader := bufio.NewReader(fd)

	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Printf("error reading file %s", err)
			break
		}
		fmt.Println(line)
	}

	return nil
}

func main() {
	flag.Parse()

	if len(flag.Args()) == 0 {
		fmt.Printf("usage: lBdkljv \n")
		return
	}

	for _, file := range flag.Args() {
		err := LineByLine(file)
		if err != nil {
			fmt.Println(err)
		}
	}
}

// pc@hp:~/P/Golang/fileHandling/readOrWriteFIleLineBYLine$ go run main.go
// usage: lBdkljv
