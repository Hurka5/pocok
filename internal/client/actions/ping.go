package actions

import (
  "net"
)

func Ping(conn net.Conn, _ []string) {
	conn.Write([]byte("PONG"))
}
