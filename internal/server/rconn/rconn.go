package rconn

import (
  "bufio"
  "os"
  "net"
  "log"
)

type RemoteConnection struct {
  conn net.Conn
  Ip   string
}

func New(c net.Conn) *RemoteConnection {
  return &RemoteConnection {
    conn: c,
    Ip:   c.RemoteAddr().String(),
  }
}

func (rc *RemoteConnection) Start() {
  go rc.loop()
}


func (rc *RemoteConnection) Send(msg string) {

  // Sending data
	_, err := rc.conn.Write([]byte(msg))
	if err != nil {
		log.Println("Error writing to connection: ", err.Error())
		return
  }

}

func (rc *RemoteConnection) Recieve() {
  
  buf := make([]byte, 1024)

	_, err := rc.conn.Read(buf)
	if err != nil {
		log.Println("Error reading from connection: ", err.Error())
		return
	}

	str := string(buf)
	log.Println("Received from client:", str)
}

func (rc *RemoteConnection) RecieveFile(path string) {

	log.Println("Receiving screenshot...")

	file, err := os.Create(path)
	if err != nil {
		log.Println("Error creating file:", err.Error())
		return
	}
	defer file.Close()

	r := bufio.NewReader(rc.conn)

	_, err = r.WriteTo(file)

	log.Println("Screenshot saved to ", path)
}

func (rc *RemoteConnection) IsAlive() bool {
  return rc.conn != nil
}

// -- private --

func (rc *RemoteConnection) loop() {
  for rc.IsAlive() {
    
  } 
}
