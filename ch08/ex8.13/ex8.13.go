// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"
)

type client chan<- string // an outgoing message channel

var (
	entering             = make(chan client)
	leaving              = make(chan client)
	messages             = make(chan string) // all incoming client messages
	registeringAddress   = make(chan string)
	unregisteringAddress = make(chan string)
)

func broadcaster() {
	clients := make(map[client]bool) // all connected clients
	addresses := make(map[string]bool)
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for cli := range clients {
				cli <- msg
			}

		case cli := <-entering:
			clients[cli] = true

		case cli := <-leaving:
			delete(clients, cli)
			close(cli)

		case address := <-registeringAddress:
			addresses[address] = true
			allClients := "All clients:"
			for addr := range addresses {
				allClients = fmt.Sprintf("%s\n%s", allClients, addr)
			}
			go func() { messages <- allClients }() // Must call a new goroutine to avoid blocking
		case address := <-unregisteringAddress:
			delete(addresses, address)
			allClients := "All clients:"
			for addr := range addresses {
				allClients = fmt.Sprintf("%s\n%s", allClients, addr)
			}
			go func() { messages <- allClients }() // Must call a new goroutine to avoid blocking
		}

	}
}

func countIdleTime(conn net.Conn, notIdleCh <-chan bool) {
	ticker := time.NewTicker(time.Second)
	counter := 0
	max := 20 // 20 seconds
	for {
		select {
		case <-ticker.C:
			counter++
			if counter == max {
				msg := conn.RemoteAddr().String() + " idle too long. Kicked out."
				messages <- msg
				fmt.Fprintln(conn, msg) // Let to-be-closed client see this msg
				ticker.Stop()
				conn.Close()
				return
			}
		case <-notIdleCh:
			counter = 0
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // outgoing client messages
	go clientWriter(conn, ch)

	who := conn.RemoteAddr().String()
	ch <- "You are " + who
	messages <- who + " has arrived"
	entering <- ch
	registeringAddress <- who

	notIdleCh := make(chan bool)
	go countIdleTime(conn, notIdleCh)

	input := bufio.NewScanner(conn)
	for input.Scan() {
		notIdleCh <- true
		messages <- who + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- ch
	messages <- who + " has left"
	unregisteringAddress <- who
	conn.Close()
}

func clientWriter(conn net.Conn, ch <-chan string) {
	for msg := range ch {
		fmt.Fprintln(conn, msg) // NOTE: ignoring network errors
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	go broadcaster()
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
