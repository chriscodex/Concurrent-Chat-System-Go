package main

import (
	"bufio"
	"flag"
	"fmt"
	"net"
)

type client chan<- string

var (
	incomingClients = make(chan client)
	leavingClients  = make(chan client)

	// Global channel
	messagesClients = make(chan string)
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port")
)

// Function to handle the connection of clients with the server
func handleConnection(conn net.Conn) {
	defer conn.Close()

	// Messages from the client that has connected
	messageChat := make(chan string)
	go messageWrite(conn, messageChat)

	// Remote Network Address
	clientName := conn.RemoteAddr().String()

	messageChat <- fmt.Sprintf("Welcome to the server, your name %s\n", clientName)

	// Message for all clients
	messagesClients <- fmt.Sprintf("New client is here, name %s\n", clientName)

	incomingClients <- messageChat

	inputMessage := bufio.NewScanner(conn)
	for inputMessage.Scan() {
		messagesClients <- fmt.Sprintf("%s: %s\n", clientName, inputMessage.Text())
	}
	leavingClients <- messageChat
	messagesClients <- fmt.Sprintf("Client %s left the chat\n", clientName)
}

func messageWrite(conn net.Conn, messages <-chan string) {
	for message := range messages {
		fmt.Fprintln(conn, message)
	}
}

func broadcast() {
	clients := make(map[client]bool)
	for {
		select {
		case message := <-messagesClients:
			for client := range clients {
				client <- message
			}
		case newClient := <-incomingClients:
			clients[newClient] = true
		case clientLeft := <-leavingClients:
			delete(clients, clientLeft)
			close(leavingClients)
		}
	}
}
