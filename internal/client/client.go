package client

import (
  "log"
  "net"
  "fmt"
)

const (
	HOST = "localhost"
	PORT = 7019
)

func Run() {
	log.Println("Initializing client...")

	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", HOST, PORT))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to server")

	HandleConnection(conn)
}
