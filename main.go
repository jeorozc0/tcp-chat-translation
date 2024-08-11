package main

import (
	"log"
	"net"
)

func main() {
	s := newServer()
	go s.run()

	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("unable to start server: %s", err.Error())
	}

	defer listener.Close()
	log.Printf("started Server on :8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Printf("Unable to accept connection: %s", err.Error())
			continue
		}
		go s.neClient(conn)
	}
}
