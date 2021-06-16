package main

import (
	"github.com/mstzn/modanisa_backend/server"
)

func main() {
	newServer := server.Server{
		Port: 3000,
	}

	go newServer.Start()

}
