package main

import (
	"flag"
	"net"
)

type client chan<- string

var (
	incomingClients = make(chan client)
	leavingClients  = make(chan client)
	messagesClients = make(chan string)
)

var (
	host = flag.String("h", "localhost", "host")
	port = flag.Int("p", 3090, "port")
)

// Function to handle the connection of clients with the server
func handleConnection(conn net.Conn) {
	defer conn.Close()
	messageChat := make(chan string)

}
