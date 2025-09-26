package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
)

func handler(conn net.Conn, ch chan string) {
	defer conn.Close()

	reader := bufio.NewReader(conn)

	for {
		str, err := reader.ReadString('\n')

		if err != nil {
			if err != io.EOF {
				fmt.Println("Server: Failed to read data, err =", err)
			}
			break
		}

		fmt.Printf("Server: Received from client msg = %s", str)

		ch <- conn.RemoteAddr().String()

		conn.Write([]byte("ok"))
	}

}

func logger(ch chan string) {
	for {
		x := <-ch
		fmt.Printf("Logger: Logging client address = %v\n\n", x)
	}
}

func server(listener net.Listener, ch chan string) {
	for {
		c, err := listener.Accept()
		if err != nil {
			continue
		}

		go handler(c, ch)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":5000")
	if err != nil {
		panic(err)
	}

	fmt.Println("Listening on localhost:5000...")
	fmt.Println("Send something using ncat: echo hello | ncat localhost 5000\n")

	ch := make(chan string)

	go logger(ch)
	server(listener, ch)
}
