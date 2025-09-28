package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	WORKERS = 5
	TASKS   = 10
)

func main() {
	var wg sync.WaitGroup

	wg.Add(WORKERS)

	go pool(&wg, WORKERS, TASKS)

	wg.Wait()

	fmt.Printf("\nFinished all tasks!")
}

func pool(wg *sync.WaitGroup, workers int, tasks int) {
	ch_result := make(chan int)
	ch_id := make(chan int)

	fmt.Printf("Creating worker pool of %v...\n\n", workers)

	for i := 1; i <= workers; i++ {
		go worker(i, ch_result, ch_id, wg)
	}

	fmt.Printf("Assigning %v tasks among %v workers...\n\n", tasks, workers)
	for i := 1; i <= tasks; i++ {
		task(i, ch_result, ch_id)
	}

	close(ch_result)
	close(ch_id)
}

func task(input int, ch_result chan int, ch_id chan int) {
	output := input * 2
	time.Sleep(5 * time.Millisecond)
	ch_result <- output
	ch_id <- input
}

func worker(worker_id int, ch_result <-chan int, ch_id <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		result, ok := <-ch_result
		task := <-ch_id

		if !ok {
			return
		}

		fmt.Printf("Worker %v, task %v, result %v\n", worker_id, task, result)
	}
}
