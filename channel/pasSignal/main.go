package main

import (
	"fmt"
	"time"
)

//Routine 1: Starting Work..
//Routine 2: Work Done..

func main() {
	starCh := make(chan struct{})
	doneCh := make(chan struct{})

	go func() {
		for {
			<-doneCh
			fmt.Println("Routine 1: Starting Work..")
			starCh <- struct{}{}
		}
	}()

	go func() {
		for {
			<-starCh
			fmt.Println("Routine 2: Work Done...")
			doneCh <- struct{}{}
		}
	}()
	doneCh <- struct{}{}
	time.Sleep(1 * time.Second)
}
