package main

import (
	"fmt"
	"time"
)

func tick(d time.Duration) <-chan int {
	c := make(chan int)
	go func() {
		time.Sleep(d)
		c <- 1
	}()
	return c
}

func main() {
	for i := 0; i < 10; i++ {
		c := tick(100 * time.Millisecond)
		x := <-c
		fmt.Printf("Main: Timer = %v, Value = %v\n", i, x)
	}
}
