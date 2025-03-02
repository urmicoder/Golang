package main

import (
	"fmt"
	"sync"
)

func resiver(workerId int, taskCh <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for task := range taskCh {
		fmt.Printf("Worker %d Proccessing task %d\n", workerId, task)
	}
}

func main() {
	const worker = 3
	const tasks = 5
	taskCh := make(chan int, tasks)
	wg := sync.WaitGroup{}

	for i := 0; i < worker; i++ {
		wg.Add(1)
		go resiver(i, taskCh, &wg)

	}

	//send data
	for i := 0; i < tasks; i++ {
		taskCh <- i
	}
	close(taskCh)
	wg.Wait()
}
