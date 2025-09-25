package main

import (
	"fmt"
	"time"
)

func main() {
	var Ball int
	table := make(chan int)

	go player(table, 1)
	go player(table, 2)

	fmt.Println("Referee, Initial Rally = ", Ball, "\n")

	table <- Ball
	time.Sleep(1 * time.Second)

	final_value := <-table
	fmt.Println("\nReferee, Final Rally = ", final_value)
}

func player(table chan int, player_num int) {
	for {
		ball := <-table
		ball++
		time.Sleep(100 * time.Millisecond)
		table <- ball
		fmt.Printf("Player %v Hits, Rally = %v\n", player_num, ball)
	}
}
