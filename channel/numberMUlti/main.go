package main

import (
	"fmt"
	"sync"
)

func sender(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	n := 5
	for i := 1; i <= 10; i++ {
		ch <- i * n
	}
	close(ch)
}

func resiver(ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range ch {
		fmt.Println(data)
	}
}

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go sender(ch, &wg)
	go resiver(ch, &wg)
	wg.Wait()
}
