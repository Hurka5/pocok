package server

import (
	"fmt"
  "strconv"
)

type Action func(*PocokServer, []string)

var OfflineActions = map[string]Action {
  "list":    List,
  "connect": Connect,
}

var OnlineActions = map[string]Action {
  "screen": Screen,
}


func Connect(ps *PocokServer, args []string) {

  index, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Println(args[0]," is not a index number")
		return
	}

  if index < 0 || len(ps.Conns) <= index {
		fmt.Println(args[0]," is out of range")
		return
  }

  ps.Curr = index

  ps.RCIp = ps.Current().Ip + " "
}


func Exit(ps *PocokServer, args []string) {
  ps.Curr = NO_CONNECTION
  ps.RCIp = ""
}


func List(ps *PocokServer, args []string) {

  fmt.Println("List of current connection:")

  for i, c := range ps.Conns {
     fmt.Println("\t", i, " ", c.Ip)
  }

}


func Screen(ps *PocokServer, args []string) {

  filename := args[0]

  if filename == "" {
    filename = "screenshot.png"
  }

  ps.Current().RecieveFile(filename)
}
