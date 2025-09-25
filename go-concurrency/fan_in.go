package main

import (
	"fmt"
	"time"
)

type Package struct {
	num   int
	value int
}

func producer(ch chan Package, d time.Duration, producer_num int) {
	for i := 0; i < 10; i++ {
		p1 := Package{num: producer_num, value: int(i)}
		ch <- p1
		time.Sleep(d)
	}
}

func reader(out chan Package) {
	for i := 0; i < 20; i++ {
		<-out
		fmt.Println("\n----------------------------------")
		fmt.Println("Reader: Received package from Main")
		fmt.Println("----------------------------------\n")
	}
}

func main() {
	ch := make(chan Package)
	out := make(chan Package)
	go producer(ch, 10*time.Millisecond, 1)
	go producer(ch, 25*time.Millisecond, 2)
	go reader(out)
	for i := 0; i < 20; i++ {
		x := <-ch
		fmt.Printf("Main: Receiving from Producer %v, Value %v\n", x.num, x.value)
		fmt.Println("Main: Sending to Reader...\n")
		out <- x
	}
}
