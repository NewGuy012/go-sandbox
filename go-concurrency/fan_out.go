package main

import (
	"fmt"
	"sync"
	"time"
)

const (
	WORKERS = 5
	TASKS   = 50
)

func main() {
	var wg sync.WaitGroup

	wg.Add(WORKERS)

	go pool(&wg, WORKERS, TASKS)

	wg.Wait()
}

func pool(wg *sync.WaitGroup, workers, tasks int) {
	ch := make(chan int)

	for i := 1; i <= workers; i++ {
		go worker(i, ch, wg)
	}

	for i := 1; i <= tasks; i++ {
		process(ch, i)
	}

	close(ch)
}

func process(ch chan int, input int) {
	output := input + 3 - 2 - 1
	time.Sleep(5 * time.Millisecond)
	ch <- output
}

func worker(worker_id int, ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		task, ok := <-ch

		if !ok {
			return
		}

		d := time.Duration(task) * time.Millisecond
		time.Sleep(d)

		fmt.Printf("Worker %v, processing task %v\n", worker_id, task)
	}
}
