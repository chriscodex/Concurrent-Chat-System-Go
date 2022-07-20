package main

import (
	"flag"
	"fmt"
	"log"
	"net"
)

var (
	port = flag.Int("p", 3090, "port")
	host = flag.String("h", "localhost", "host")
)

func main() {
	flag.Parse()
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%p", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	done := make(chan struct{})

}
