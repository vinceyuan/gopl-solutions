// Chat is a server that lets clients chat with each other.
package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

type client struct {
	channel chan<- string // an outgoing message channel
	name    string
}

var (
	entering = make(chan *client)
	leaving  = make(chan *client)
	messages = make(chan string) // all incoming client messages
)

func broadcaster() {
	clients := make(map[string]*client) // all connected clients
	for {
		select {
		case msg := <-messages:
			// Broadcast incoming message to all
			// clients' outgoing message channels.
			for _, cli := range clients {
				cli.channel <- msg
			}

		case cli := <-entering:
			go giveAllClients(cli.channel, clients)
			clients[cli.name] = cli

		case cli := <-leaving:
			delete(clients, cli.name)
			close(cli.channel)
		}
	}
}

func giveAllClients(channel chan<- string, clients map[string]*client) {
	if len(clients) > 1 {
		channel <- "All clients:"
		for _, cli := range clients {
			channel <- cli.name
		}
	}
}

func handleConn(conn net.Conn) {
	ch := make(chan string) // Necessary because me.channel = ch will do implicit conversion and line:65 will fail
	me := &client{channel: ch, name: conn.RemoteAddr().String()}
	go clientWriter(conn, ch)

	me.channel <- "You are " + me.name
	messages <- me.name + " has arrived"
	entering <- me

	input := bufio.NewScanner(conn)
	for input.Scan() {
		messages <- me.name + ": " + input.Text()
	}
	// NOTE: ignoring potential errors from input.Err()

	leaving <- me
	messages <- me.name + " has left"
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
