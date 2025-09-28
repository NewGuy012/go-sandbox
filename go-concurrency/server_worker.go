package main

import (
	"fmt"
	"net"
	"time"
)

const (
	WORKERS = 5
)

func pool(ch chan int, workers int) {
	wch := make(chan int)
	results := make(chan int)

	fmt.Printf("Creating worker pool of %v...\n\n", workers)

	for i := 0; i < workers; i++ {
		go worker(wch, results)
	}

	go parse(results)

	for {
		val := <-ch
		wch <- val
	}
}

func worker(wch chan int, results chan int) {
	for {
		data := <-wch
		data++
		results <- data
	}
}

func parse(results chan int) {
	for {
		<-results
	}
}

func server(l net.Listener, ch chan int) {
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}

		go handler(c, ch)
	}
}

func handler(conn net.Conn, ch chan int) {
	defer conn.Close()

	time.Sleep(50 * time.Millisecond)
	ch <- 1
	conn.Write([]byte("ok"))
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on localhost:5000...")
	fmt.Println("Send something using ncat: echo hello | ncat localhost 5000\n")

	ch := make(chan int)

	go pool(ch, WORKERS)
	go server(listener, ch)
}
