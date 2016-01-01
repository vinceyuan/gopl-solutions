package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func handleConn(c net.Conn) {
	counter := make(chan int, 1) // Used a buffered channel to avoid go routine leak

	input := bufio.NewScanner(c)
	go func() {
		for input.Scan() {
			counter <- 10 // Reset counter
			go echo(c, input.Text(), 1*time.Second)
		}
		counter <- 0
		fmt.Println("input.Scan() go routine ends")
	}()

	for i := 10; i >= 0; {
		fmt.Println(i)
		select {
		case value := <-counter:
			i = value

		default:
			i--
		}
		time.Sleep(time.Second)
	}
	c.Close()
	fmt.Println("Closed a connection")

}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
