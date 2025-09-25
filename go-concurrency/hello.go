package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		ch <- "Hello World!"
	}()

	x := <-ch
	fmt.Println("Main:", x)
}
