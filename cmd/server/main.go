package main

import (
  "pocok/internal/server"
)

func main() {

  // Initialize server
  srv := server.New()

  // Start listener
  srv.Listen()

  // Start SHell
  srv.Shell()
}
